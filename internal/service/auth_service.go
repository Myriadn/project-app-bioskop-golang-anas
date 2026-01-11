package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/config"
	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/repository"
	"project-app-bioskop-golang-homework-anas/internal/utils"

	"go.uber.org/zap"
)

type AuthService interface {
	Register(ctx context.Context, req *domain.RegisterRequest) (*domain.AuthResponse, error)
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error)
	Logout(ctx context.Context, token string) error
	ValidateToken(ctx context.Context, token string) (*domain.User, error)
}

type authService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.AuthTokenRepository
	config    *config.Config
	logger    *zap.Logger
}

func NewAuthService(
	userRepo repository.UserRepository,
	tokenRepo repository.AuthTokenRepository,
	config *config.Config,
	logger *zap.Logger,
) AuthService {
	return &authService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		config:    config,
		logger:    logger,
	}
}

func (s *authService) Register(ctx context.Context, req *domain.RegisterRequest) (*domain.AuthResponse, error) {
	// Check if username already exists
	existingUser, _ := s.userRepo.GetByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Check if email already exists
	existingUser, _ = s.userRepo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("Failed to hash password", zap.Error(err))
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &domain.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		IsVerified:   false, // Default false, bisa di-set true jika tidak pakai email verification
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	s.logger.Info("User registered successfully",
		zap.Int("user_id", user.ID),
		zap.String("username", user.Username),
	)

	// Generate token
	tokenString, err := s.generateToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		User:  user,
		Token: tokenString,
	}, nil
}

func (s *authService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error) {
	// Get user by username
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("invalid username or password")
		}
		s.logger.Error("Failed to get user", zap.Error(err))
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid username or password")
	}

	s.logger.Info("User logged in successfully",
		zap.Int("user_id", user.ID),
		zap.String("username", user.Username),
	)

	// Generate token
	tokenString, err := s.generateToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		User:  user,
		Token: tokenString,
	}, nil
}

func (s *authService) Logout(ctx context.Context, token string) error {
	// Delete token
	if err := s.tokenRepo.Delete(ctx, token); err != nil {
		s.logger.Error("Failed to delete token", zap.Error(err))
		return fmt.Errorf("failed to logout: %w", err)
	}

	s.logger.Info("User logged out successfully")
	return nil
}

func (s *authService) ValidateToken(ctx context.Context, token string) (*domain.User, error) {
	// Get token from database
	authToken, err := s.tokenRepo.GetByToken(ctx, token)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	// Get user
	user, err := s.userRepo.GetByID(ctx, authToken.UserID)
	if err != nil {
		s.logger.Error("Failed to get user", zap.Error(err))
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *authService) generateToken(ctx context.Context, userID int) (string, error) {
	// Generate random token
	tokenString, err := utils.GenerateToken(32)
	if err != nil {
		s.logger.Error("Failed to generate token", zap.Error(err))
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	// Create auth token record
	authToken := &domain.AuthToken{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(s.config.Token.ExpiryTime),
	}

	if err := s.tokenRepo.Create(ctx, authToken); err != nil {
		s.logger.Error("Failed to save token", zap.Error(err))
		return "", fmt.Errorf("failed to save token: %w", err)
	}

	return tokenString, nil
}
