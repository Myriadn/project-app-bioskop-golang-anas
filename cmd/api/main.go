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
	logger.Info("Repositories initialized")

	// Initialize Services
	authService := service.NewAuthService(userRepo, authTokenRepo, cfg, logger.Log)
	cinemaService := service.NewCinemaService(cinemaRepo, logger.Log)
	seatService := service.NewSeatService(seatRepo, showtimeRepo, cinemaRepo, logger.Log)
	paymentMethodService := service.NewPaymentMethodService(paymentMethodRepo, logger.Log)
	logger.Info("Services initialized")

	// Initialize Handlers
	authHandler := handler.NewAuthHandler(authService, logger.Log)
	cinemaHandler := handler.NewCinemaHandler(cinemaService, logger.Log)
	seatHandler := handler.NewSeatHandler(seatService, logger.Log)
	paymentMethodHandler := handler.NewPaymentMethodHandler(paymentMethodService, logger.Log)
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
		fmt.Printf("\n🚀 Server is running on http://localhost%s\n", server.Addr)
		fmt.Printf("📚 API Endpoints:\n")
		fmt.Printf("   GET  http://localhost%s/health\n", server.Addr)
		fmt.Printf("   POST http://localhost%s/api/register\n", server.Addr)
		fmt.Printf("   POST http://localhost%s/api/login\n", server.Addr)
		fmt.Printf("   POST http://localhost%s/api/logout (Protected)\n", server.Addr)
		fmt.Printf("   GET  http://localhost%s/api/cinemas\n", server.Addr)
		fmt.Printf("   GET  http://localhost%s/api/cinemas/{id}\n", server.Addr)
		fmt.Printf("   GET  http://localhost%s/api/cinemas/{id}/seats?date=YYYY-MM-DD&time=HH:MM:SS\n", server.Addr)
		fmt.Printf("   GET  http://localhost%s/api/payment-methods\n\n", server.Addr)

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
