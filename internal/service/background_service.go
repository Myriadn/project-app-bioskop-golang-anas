package service

import (
	"context"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/repository"

	"go.uber.org/zap"
)

type BackgroundService interface {
	StartTokenCleanup(interval time.Duration)
	StartOTPCleanup(interval time.Duration)
	Stop()
}

type backgroundService struct {
	tokenRepo repository.AuthTokenRepository
	otpRepo   repository.OTPRepository
	logger    *zap.Logger
	stopChan  chan bool
}

func NewBackgroundService(
	tokenRepo repository.AuthTokenRepository,
	otpRepo repository.OTPRepository,
	logger *zap.Logger,
) BackgroundService {
	return &backgroundService{
		tokenRepo: tokenRepo,
		otpRepo:   otpRepo,
		logger:    logger,
		stopChan:  make(chan bool),
	}
}

// StartTokenCleanup menjalankan background job untuk cleanup expired tokens
func (s *backgroundService) StartTokenCleanup(interval time.Duration) {
	s.logger.Info("Starting token cleanup background job", zap.Duration("interval", interval))

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

// StartOTPCleanup menjalankan background job untuk cleanup expired OTPs
func (s *backgroundService) StartOTPCleanup(interval time.Duration) {
	s.logger.Info("Starting OTP cleanup background job", zap.Duration("interval", interval))

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.cleanupExpiredOTPs()
			case <-s.stopChan:
				s.logger.Info("OTP cleanup background job stopped")
				return
			}
		}
	}()
}

func (s *backgroundService) cleanupExpiredTokens() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

func (s *backgroundService) cleanupExpiredOTPs() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		s.logger.Info("Running OTP cleanup...")

		err := s.otpRepo.DeleteExpired(ctx)
		if err != nil {
			s.logger.Error("Failed to cleanup expired OTPs", zap.Error(err))
			return
		}

		s.logger.Info("OTP cleanup completed successfully")
	}()
}

func (s *backgroundService) Stop() {
	s.stopChan <- true
}
