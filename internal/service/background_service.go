package service

import (
	"context"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/repository"

	"go.uber.org/zap"
)

type BackgroundService interface {
	StartTokenCleanup(interval time.Duration)
	Stop()
}

type backgroundService struct {
	tokenRepo repository.AuthTokenRepository
	logger    *zap.Logger
	stopChan  chan bool
}

func NewBackgroundService(
	tokenRepo repository.AuthTokenRepository,
	logger *zap.Logger,
) BackgroundService {
	return &backgroundService{
		tokenRepo: tokenRepo,
		logger:    logger,
		stopChan:  make(chan bool),
	}
}

// StartTokenCleanup menjalankan background job untuk cleanup expired tokens
func (s *backgroundService) StartTokenCleanup(interval time.Duration) {
	s.logger.Info("Starting token cleanup background job", zap.Duration("interval", interval))

	// Run in goroutine
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.cleanupExpiredTokens()
			case <-s.stopChan:
				s.logger.Info("Token cleanup background job stopped")
				return
			}
		}
	}()
}

func (s *backgroundService) cleanupExpiredTokens() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Run cleanup in goroutine
	go func() {
		s.logger.Info("Running token cleanup...")

		err := s.tokenRepo.DeleteExpired(ctx)
		if err != nil {
			s.logger.Error("Failed to cleanup expired tokens", zap.Error(err))
			return
		}

		s.logger.Info("Token cleanup completed successfully")
	}()
}

func (s *backgroundService) Stop() {
	s.stopChan <- true
}
