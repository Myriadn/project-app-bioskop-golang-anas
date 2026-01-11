package service

import (
	"context"
	"fmt"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/repository"
	"project-app-bioskop-golang-homework-anas/internal/utils"

	"go.uber.org/zap"
)

type OTPService interface {
	SendOTP(ctx context.Context, userID int, email, username string) error
	VerifyOTP(ctx context.Context, email, code string) error
	ResendOTP(ctx context.Context, email string) error
}

type otpService struct {
	otpRepo      repository.OTPRepository
	userRepo     repository.UserRepository
	emailService *utils.EmailService
	logger       *zap.Logger
}

func NewOTPService(
	otpRepo repository.OTPRepository,
	userRepo repository.UserRepository,
	emailService *utils.EmailService,
	logger *zap.Logger,
) OTPService {
	return &otpService{
		otpRepo:      otpRepo,
		userRepo:     userRepo,
		emailService: emailService,
		logger:       logger,
	}
}

func (s *otpService) SendOTP(ctx context.Context, userID int, email, username string) error {
	// Generate OTP
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		s.logger.Error("Failed to generate OTP", zap.Error(err))
		return fmt.Errorf("failed to generate OTP")
	}

	// Delete old OTPs for this user
	if err := s.otpRepo.DeleteByUserID(ctx, userID); err != nil {
		s.logger.Warn("Failed to delete old OTPs", zap.Error(err))
	}

	// Create OTP record (expires in 10 minutes)
	otp := &domain.OTPCode{
		UserID:    userID,
		Code:      otpCode,
		IsUsed:    false,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	if err := s.otpRepo.Create(ctx, otp); err != nil {
		s.logger.Error("Failed to save OTP", zap.Error(err))
		return fmt.Errorf("failed to save OTP")
	}

	// Send email async (goroutine)
	s.emailService.SendEmailAsync(email, username, otpCode)

	s.logger.Info("OTP sent successfully",
		zap.Int("user_id", userID),
		zap.String("email", email),
	)

	return nil
}

func (s *otpService) VerifyOTP(ctx context.Context, email, code string) error {
	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Check if already verified
	if user.IsVerified {
		return fmt.Errorf("account already verified")
	}

	// Validate OTP
	otp, err := s.otpRepo.GetByUserIDAndCode(ctx, user.ID, code)
	if err != nil {
		s.logger.Error("Invalid OTP", zap.Error(err))
		return fmt.Errorf("invalid or expired OTP code")
	}

	// Mark OTP as used
	if err := s.otpRepo.MarkAsUsed(ctx, otp.ID); err != nil {
		s.logger.Error("Failed to mark OTP as used", zap.Error(err))
	}

	// Update user verification status
	user.IsVerified = true
	if err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.Error("Failed to verify user", zap.Error(err))
		return fmt.Errorf("failed to verify account")
	}

	s.logger.Info("User verified successfully",
		zap.Int("user_id", user.ID),
		zap.String("email", email),
	)

	// Send welcome email async
	s.emailService.SendWelcomeEmailAsync(email, user.Username)

	return nil
}

func (s *otpService) ResendOTP(ctx context.Context, email string) error {
	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Check if already verified
	if user.IsVerified {
		return fmt.Errorf("account already verified")
	}

	// Send new OTP
	return s.SendOTP(ctx, user.ID, email, user.Username)
}
