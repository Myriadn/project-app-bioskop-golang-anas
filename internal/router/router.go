package router

import (
	"net/http"

	"project-app-bioskop-golang-homework-anas/internal/handler"
	"project-app-bioskop-golang-homework-anas/internal/middleware"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Router struct {
	authHandler          *handler.AuthHandler
	cinemaHandler        *handler.CinemaHandler
	seatHandler          *handler.SeatHandler
	paymentMethodHandler *handler.PaymentMethodHandler
	authMiddleware       *middleware.AuthMiddleware
	logger               *zap.Logger
}

func NewRouter(
	authHandler *handler.AuthHandler,
	cinemaHandler *handler.CinemaHandler,
	seatHandler *handler.SeatHandler,
	paymentMethodHandler *handler.PaymentMethodHandler,
	authMiddleware *middleware.AuthMiddleware,
	logger *zap.Logger,
) *Router {
	return &Router{
		authHandler:          authHandler,
		cinemaHandler:        cinemaHandler,
		seatHandler:          seatHandler,
		paymentMethodHandler: paymentMethodHandler,
		authMiddleware:       authMiddleware,
		logger:               logger,
	}
}

func (rt *Router) SetupRoutes() http.Handler {
	r := chi.NewRouter()

	// Global middlewares
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.RequestID)
	r.Use(middleware.CORSMiddleware)
	r.Use(middleware.LoggerMiddleware(rt.logger))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Cinema Booking API is running"}`))
	})

	// API routes
	r.Route("/api", func(r chi.Router) {
		// Auth routes (public)
		rt.setupAuthRoutes(r)

		// Cinema routes (public)
		rt.setupCinemaRoutes(r)

		// Payment method routes (public)
		rt.setupPaymentMethodRoutes(r)

		// Protected routes (auth required)
		r.Group(func(r chi.Router) {
			r.Use(rt.authMiddleware.RequireAuth)

			// Logout
			r.Post("/logout", rt.authHandler.Logout)

			// TODO: Booking routes (akan dibuat selanjutnya)
			// TODO: User booking history routes (akan dibuat selanjutnya)
		})
	})

	return r
}

// setupAuthRoutes mengatur routing untuk authentication
func (rt *Router) setupAuthRoutes(r chi.Router) {
	r.Post("/register", rt.authHandler.Register)
	r.Post("/login", rt.authHandler.Login)
}

// setupCinemaRoutes mengatur routing untuk cinema
func (rt *Router) setupCinemaRoutes(r chi.Router) {
	r.Get("/cinemas", rt.cinemaHandler.GetAllCinemas)
	r.Get("/cinemas/{cinemaId}", rt.cinemaHandler.GetCinemaByID)
	r.Get("/cinemas/{cinemaId}/seats", rt.seatHandler.GetSeatAvailability)
}

// setupPaymentMethodRoutes mengatur routing untuk payment methods
func (rt *Router) setupPaymentMethodRoutes(r chi.Router) {
	r.Get("/payment-methods", rt.paymentMethodHandler.GetAllPaymentMethods)
}
