package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/config"
	"project-app-bioskop-golang-homework-anas/internal/handler"
	"project-app-bioskop-golang-homework-anas/internal/middleware"
	"project-app-bioskop-golang-homework-anas/internal/repository"
	"project-app-bioskop-golang-homework-anas/internal/router"
	"project-app-bioskop-golang-homework-anas/internal/service"
	"project-app-bioskop-golang-homework-anas/pkg/database"
	"project-app-bioskop-golang-homework-anas/pkg/logger"
	"project-app-bioskop-golang-homework-anas/pkg/validator"

	"go.uber.org/zap"
)

func main() {
	// Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// Initialize Logger
	if err := logger.InitLogger(cfg.Log.Level, cfg.Log.File); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("Starting Cinema Booking API",
		zap.String("app_name", cfg.App.Name),
		zap.String("environment", cfg.App.Env),
		zap.String("port", cfg.App.Port),
	)

	// Initialize Validator
	validator.InitValidator()
	logger.Info("Validator initialized")

	// Connect to Database
	dsn := cfg.GetDatabaseDSN()
	db, err := database.NewPostgresPool(dsn, logger.Log)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer database.ClosePool(db, logger.Log)

	// Initialize Repositories
	userRepo := repository.NewUserRepository(db)
	authTokenRepo := repository.NewAuthTokenRepository(db)
	cinemaRepo := repository.NewCinemaRepository(db)
	showtimeRepo := repository.NewShowtimeRepository(db)
	seatRepo := repository.NewSeatRepository(db)
	paymentMethodRepo := repository.NewPaymentMethodRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	logger.Info("Repositories initialized")

	// Initialize Services
	authService := service.NewAuthService(userRepo, authTokenRepo, cfg, logger.Log)
	cinemaService := service.NewCinemaService(cinemaRepo, logger.Log)
	seatService := service.NewSeatService(seatRepo, showtimeRepo, cinemaRepo, logger.Log)
	paymentMethodService := service.NewPaymentMethodService(paymentMethodRepo, logger.Log)
	bookingService := service.NewBookingService(bookingRepo, showtimeRepo, seatRepo, paymentMethodRepo, logger.Log)
	paymentService := service.NewPaymentService(paymentRepo, bookingRepo, paymentMethodRepo, logger.Log)
	logger.Info("Services initialized")

	// Initialize Handlers
	authHandler := handler.NewAuthHandler(authService, logger.Log)
	cinemaHandler := handler.NewCinemaHandler(cinemaService, logger.Log)
	seatHandler := handler.NewSeatHandler(seatService, logger.Log)
	paymentMethodHandler := handler.NewPaymentMethodHandler(paymentMethodService, logger.Log)
	bookingHandler := handler.NewBookingHandler(bookingService, logger.Log)
	paymentHandler := handler.NewPaymentHandler(paymentService, logger.Log)
	logger.Info("Handlers initialized")

	// Initialize Middlewares
	authMiddleware := middleware.NewAuthMiddleware(authService, logger.Log)
	logger.Info("Middlewares initialized")

	// Setup Router
	appRouter := router.NewRouter(
		authHandler,
		cinemaHandler,
		seatHandler,
		paymentMethodHandler,
		bookingHandler,
		paymentHandler,
		authMiddleware,
		logger.Log,
	)
	httpHandler := appRouter.SetupRoutes()
	logger.Info("Router configured")

	// Create HTTP Server
	server := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      httpHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		logger.Info("Server starting", zap.String("address", server.Addr))
		fmt.Printf("\n🚀 Cinema Booking API is running on http://localhost%s\n\n", server.Addr)
		fmt.Printf("📚 Available Endpoints:\n")
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
		fmt.Printf("🔓 PUBLIC ENDPOINTS:\n")
		fmt.Printf("   GET  /health                          - Health check\n")
		fmt.Printf("   POST /api/register                    - Register user\n")
		fmt.Printf("   POST /api/login                       - Login user\n")
		fmt.Printf("   GET  /api/cinemas                     - Get all cinemas\n")
		fmt.Printf("   GET  /api/cinemas/{id}                - Get cinema detail\n")
		fmt.Printf("   GET  /api/cinemas/{id}/seats          - Get seat availability\n")
		fmt.Printf("   GET  /api/payment-methods             - Get payment methods\n")
		fmt.Printf("   POST /api/pay                         - Process payment\n")
		fmt.Printf("\n🔒 PROTECTED ENDPOINTS (Require Token):\n")
		fmt.Printf("   POST /api/logout                      - Logout user\n")
		fmt.Printf("   POST /api/booking                     - Create booking\n")
		fmt.Printf("   GET  /api/user/bookings               - Get user bookings\n")
		fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("Server shutting down...")
	fmt.Println("\n🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exited gracefully")
	fmt.Println("✅ Server stopped")
}
