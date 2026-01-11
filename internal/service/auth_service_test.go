package service

import (
	"context"
	"errors"
	"testing"

	"project-app-bioskop-golang-homework-anas/internal/config"
	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// Mock repositories
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

type MockAuthTokenRepository struct {
	mock.Mock
}

func (m *MockAuthTokenRepository) Create(ctx context.Context, token *domain.AuthToken) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

func (m *MockAuthTokenRepository) GetByToken(ctx context.Context, token string) (*domain.AuthToken, error) {
	args := m.Called(ctx, token)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.AuthToken), args.Error(1)
}

func (m *MockAuthTokenRepository) Delete(ctx context.Context, token string) error {
	args := m.Called(ctx, token)
	return args.Error(0)
}

func (m *MockAuthTokenRepository) DeleteByUserID(ctx context.Context, userID int) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func (m *MockAuthTokenRepository) DeleteExpired(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

type MockOTPService struct {
	mock.Mock
}

func (m *MockOTPService) SendOTP(ctx context.Context, userID int, email, username string) error {
	args := m.Called(ctx, userID, email, username)
	return args.Error(0)
}

func (m *MockOTPService) VerifyOTP(ctx context.Context, email, code string) error {
	args := m.Called(ctx, email, code)
	return args.Error(0)
}

func (m *MockOTPService) ResendOTP(ctx context.Context, email string) error {
	args := m.Called(ctx, email)
	return args.Error(0)
}

func TestAuthService_Register_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockTokenRepo := new(MockAuthTokenRepository)
	mockOTPService := new(MockOTPService)

	cfg := &config.Config{
		Token: config.TokenConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}

	logger, _ := zap.NewDevelopment()
	authService := NewAuthService(mockUserRepo, mockTokenRepo, mockOTPService, cfg, logger)

	req := &domain.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Mock expectations
	mockUserRepo.On("GetByUsername", mock.Anything, req.Username).Return(nil, errors.New("not found"))
	mockUserRepo.On("GetByEmail", mock.Anything, req.Email).Return(nil, errors.New("not found"))
	mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)
	mockOTPService.On("SendOTP", mock.Anything, mock.AnythingOfType("int"), req.Email, req.Username).Return(nil)
	mockTokenRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.AuthToken")).Return(nil)

	// Execute
	result, err := authService.Register(context.Background(), req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Username, result.User.Username)
	assert.Equal(t, req.Email, result.User.Email)
	assert.False(t, result.User.IsVerified)
	assert.NotEmpty(t, result.Token)

	mockUserRepo.AssertExpectations(t)
	mockTokenRepo.AssertExpectations(t)
	mockOTPService.AssertExpectations(t)
}

func TestAuthService_Register_UsernameExists(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockTokenRepo := new(MockAuthTokenRepository)
	mockOTPService := new(MockOTPService)

	cfg := &config.Config{
		Token: config.TokenConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}

	logger := zap.NewNop()
	service := NewAuthService(mockUserRepo, mockTokenRepo, mockOTPService, cfg, logger)

	ctx := context.Background()
	existingUser := &domain.User{ID: 1, Username: "existing"}

	req := &domain.RegisterRequest{
		Username: "existing",
		Email:    "test@example.com",
		Password: "password123",
	}

	mockUserRepo.On("GetByUsername", ctx, "existing").Return(existingUser, nil)

	result, err := service.Register(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "username already exists")
	mockUserRepo.AssertExpectations(t)
}

func TestAuthService_Register_EmailExists(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockTokenRepo := new(MockAuthTokenRepository)
	mockOTPService := new(MockOTPService)

	cfg := &config.Config{
		Token: config.TokenConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}

	logger := zap.NewNop()
	service := NewAuthService(mockUserRepo, mockTokenRepo, mockOTPService, cfg, logger)

	ctx := context.Background()
	existingUser := &domain.User{ID: 1, Email: "test@example.com"}

	req := &domain.RegisterRequest{
		Username: "newuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	mockUserRepo.On("GetByUsername", ctx, "newuser").Return(nil, errors.New("not found"))
	mockUserRepo.On("GetByEmail", ctx, "test@example.com").Return(existingUser, nil)

	result, err := service.Register(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "email already exists")
	mockUserRepo.AssertExpectations(t)
}

func TestAuthService_Logout_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockTokenRepo := new(MockAuthTokenRepository)
	mockOTPService := new(MockOTPService)

	cfg := &config.Config{
		Token: config.TokenConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}

	logger := zap.NewNop()
	service := NewAuthService(mockUserRepo, mockTokenRepo, mockOTPService, cfg, logger)

	ctx := context.Background()

	mockTokenRepo.On("Delete", ctx, "valid-token").Return(nil)

	err := service.Logout(ctx, "valid-token")

	assert.NoError(t, err)
	mockTokenRepo.AssertExpectations(t)
}

func TestAuthService_Logout_Error(t *testing.T) {
	mockUserRepo := new(MockUserRepository)
	mockTokenRepo := new(MockAuthTokenRepository)
	mockOTPService := new(MockOTPService)

	cfg := &config.Config{
		Token: config.TokenConfig{
			Secret:      "test-secret",
			ExpiryHours: 24,
		},
	}

	logger := zap.NewNop()
	service := NewAuthService(mockUserRepo, mockTokenRepo, mockOTPService, cfg, logger)

	ctx := context.Background()

	mockTokenRepo.On("Delete", ctx, "invalid-token").Return(errors.New("token not found"))

	err := service.Logout(ctx, "invalid-token")

	assert.Error(t, err)
	mockTokenRepo.AssertExpectations(t)
}
