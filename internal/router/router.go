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
	authHandler    *handler.AuthHandler
	authMiddleware *middleware.AuthMiddleware
	logger         *zap.Logger
}

func NewRouter(
	authHandler *handler.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
	logger *zap.Logger,
) *Router {
	return &Router{
		authHandler:    authHandler,
		authMiddleware: authMiddleware,
		logger:         logger,
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

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(rt.authMiddleware.RequireAuth)

			// Logout
			r.Post("/logout", rt.authHandler.Logout)

			// TODO: Nanti akan ditambahkan routes untuk:
			// - Cinema routes
			// - Booking routes
			// - Payment routes
			// - User booking history routes
		})
	})

	return r
}

// setupAuthRoutes mengatur routing untuk authentication
func (rt *Router) setupAuthRoutes(r chi.Router) {
	r.Post("/register", rt.authHandler.Register)
	r.Post("/login", rt.authHandler.Login)
}
