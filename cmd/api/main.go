package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/config"
	"project-app-bioskop-golang-homework-anas/pkg/database"
	"project-app-bioskop-golang-homework-anas/pkg/logger"
	"project-app-bioskop-golang-homework-anas/pkg/validator"

	"go.uber.org/zap"
)

func main() {
	// 1. Load Configuration
	fmt.Println("🔧 Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("❌ Failed to load config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ Configuration loaded successfully!\n")
	fmt.Printf("   App Name: %s\n", cfg.App.Name)
	fmt.Printf("   App Port: %s\n", cfg.App.Port)
	fmt.Printf("   App Env: %s\n\n", cfg.App.Env)

	// 2. Initialize Logger
	fmt.Println("📝 Initializing logger...")
	if err := logger.InitLogger(cfg.Log.Level, cfg.Log.File); err != nil {
		fmt.Printf("❌ Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()
	fmt.Println("✅ Logger initialized successfully!\n")

	// Log dengan Zap
	logger.Info("Application starting",
		zap.String("app_name", cfg.App.Name),
		zap.String("environment", cfg.App.Env),
	)

	// 3. Initialize Validator
	fmt.Println("✔️  Initializing validator...")
	validator.InitValidator()
	logger.Info("Validator initialized successfully")
	fmt.Println("✅ Validator initialized successfully!\n")

	// 4. Connect to Database
	fmt.Println("🗄️  Connecting to database...")
	dsn := cfg.GetDatabaseDSN()
	logger.Info("Connecting to database", zap.String("dsn", dsn))

	db, err := database.NewPostgresPool(dsn, logger.Log)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
		fmt.Printf("❌ Failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer database.ClosePool(db, logger.Log)
	fmt.Println("✅ Database connected successfully!\n")

	// 5. Test Database Query
	fmt.Println("🧪 Testing database query...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test query: count users
	var userCount int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&userCount)
	if err != nil {
		logger.Error("Failed to query database", zap.Error(err))
		fmt.Printf("❌ Failed to query database: %v\n", err)
	} else {
		fmt.Printf("✅ Database query successful! Users count: %d\n\n", userCount)
		logger.Info("Database query successful", zap.Int("user_count", userCount))
	}

	// Test query: count cinemas
	var cinemaCount int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM cinemas").Scan(&cinemaCount)
	if err != nil {
		logger.Error("Failed to query cinemas", zap.Error(err))
		fmt.Printf("❌ Failed to query cinemas: %v\n", err)
	} else {
		fmt.Printf("✅ Cinemas in database: %d\n", cinemaCount)
		logger.Info("Cinemas loaded", zap.Int("cinema_count", cinemaCount))
	}

	// Test query: count movies
	var movieCount int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM movies").Scan(&movieCount)
	if err != nil {
		logger.Error("Failed to query movies", zap.Error(err))
		fmt.Printf("❌ Failed to query movies: %v\n", err)
	} else {
		fmt.Printf("✅ Movies in database: %d\n", movieCount)
		logger.Info("Movies loaded", zap.Int("movie_count", movieCount))
	}

	// Test query: count seats
	var seatCount int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM seats").Scan(&seatCount)
	if err != nil {
		logger.Error("Failed to query seats", zap.Error(err))
		fmt.Printf("❌ Failed to query seats: %v\n", err)
	} else {
		fmt.Printf("✅ Seats in database: %d\n", seatCount)
		logger.Info("Seats loaded", zap.Int("seat_count", seatCount))
	}

	// Test query: count payment methods
	var paymentMethodCount int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM payment_methods").Scan(&paymentMethodCount)
	if err != nil {
		logger.Error("Failed to query payment methods", zap.Error(err))
		fmt.Printf("❌ Failed to query payment methods: %v\n", err)
	} else {
		fmt.Printf("✅ Payment methods in database: %d\n\n", paymentMethodCount)
		logger.Info("Payment methods loaded", zap.Int("payment_method_count", paymentMethodCount))
	}

	// 6. Summary
	fmt.Println("=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")
	fmt.Println("🎉 ALL SYSTEMS READY!")
	fmt.Println("=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")
	fmt.Printf("📦 Database Summary:\n")
	fmt.Printf("   - Users: %d\n", userCount)
	fmt.Printf("   - Cinemas: %d\n", cinemaCount)
	fmt.Printf("   - Movies: %d\n", movieCount)
	fmt.Printf("   - Seats: %d\n", seatCount)
	fmt.Printf("   - Payment Methods: %d\n", paymentMethodCount)
	fmt.Println("=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")
	fmt.Println()

	logger.Info("Foundation test completed successfully! Ready to build the API! 🚀")

	fmt.Println("✨ Foundation setup is complete!")
	fmt.Println("🚀 Ready to implement Auth System (register, login, logout)")
	fmt.Println()
	fmt.Println("Press Ctrl+C to exit...")
}
