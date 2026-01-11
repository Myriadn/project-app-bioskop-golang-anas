package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// NewPostgresPool membuat connection pool ke PostgreSQL
func NewPostgresPool(dsn string, logger *zap.Logger) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	// Set connection pool settings
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	// Create pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Database connection established",
		zap.String("host", config.ConnConfig.Host),
		zap.Uint16("port", config.ConnConfig.Port),
		zap.String("database", config.ConnConfig.Database),
	)

	return pool, nil
}

// ClosePool menutup database connection pool
func ClosePool(pool *pgxpool.Pool, logger *zap.Logger) {
	if pool != nil {
		pool.Close()
		logger.Info("Database connection closed")
	}
}
