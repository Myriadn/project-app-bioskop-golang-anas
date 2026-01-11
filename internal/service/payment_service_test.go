package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockPaymentRepository struct {
	mock.Mock
}

func (m *MockPaymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	args := m.Called(ctx, payment)
	return args.Error(0)
}

func (m *MockPaymentRepository) GetByBookingID(ctx context.Context, bookingID int) (*domain.Payment, error) {
	args := m.Called(ctx, bookingID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Payment), args.Error(1)
}

func (m *MockPaymentRepository) Update(ctx context.Context, payment *domain.Payment) error {
	args := m.Called(ctx, payment)
	return args.Error(0)
}

func TestPaymentService_ProcessPayment_Success(t *testing.T) {
	mockPaymentRepo := new(MockPaymentRepository)
	mockBookingRepo := new(MockBookingRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewPaymentService(mockPaymentRepo, mockBookingRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()
	now := time.Now()

	booking := &domain.Booking{
		ID:          1,
		UserID:      1,
		Status:      "pending",
		TotalPrice:  50000,
		BookingCode: "BK123",
	}

	paymentMethod := &domain.PaymentMethod{
		ID:   1,
		Code: "CREDIT_CARD",
		Name: "Credit Card",
	}

	payment := &domain.Payment{
		ID:              1,
		BookingID:       1,
		PaymentMethodID: 1,
		Amount:          50000,
		Status:          "success",
		PaidAt:          &now,
	}

	req := &domain.PaymentRequest{
		BookingID:      1,
		PaymentMethod:  "CREDIT_CARD",
		PaymentDetails: domain.PaymentDetails{"card_type": "Visa"},
	}

	mockBookingRepo.On("GetByID", ctx, 1).Return(booking, nil)
	mockPaymentMethodRepo.On("GetByCode", ctx, "CREDIT_CARD").Return(paymentMethod, nil)
	mockPaymentRepo.On("Create", ctx, mock.AnythingOfType("*domain.Payment")).Return(nil).Run(func(args mock.Arguments) {
		p := args.Get(1).(*domain.Payment)
		p.ID = 1
	})
	mockBookingRepo.On("Update", ctx, mock.AnythingOfType("*domain.Booking")).Return(nil)
	mockPaymentRepo.On("GetByBookingID", ctx, 1).Return(payment, nil)

	result, err := service.ProcessPayment(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "success", result.Status)
	mockPaymentRepo.AssertExpectations(t)
	mockBookingRepo.AssertExpectations(t)
	mockPaymentMethodRepo.AssertExpectations(t)
}

func TestPaymentService_ProcessPayment_BookingNotFound(t *testing.T) {
	mockPaymentRepo := new(MockPaymentRepository)
	mockBookingRepo := new(MockBookingRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewPaymentService(mockPaymentRepo, mockBookingRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()
	req := &domain.PaymentRequest{
		BookingID:     999,
		PaymentMethod: "CREDIT_CARD",
	}

	mockBookingRepo.On("GetByID", ctx, 999).Return(nil, errors.New("not found"))

	result, err := service.ProcessPayment(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "booking not found")
	mockBookingRepo.AssertExpectations(t)
}

func TestPaymentService_ProcessPayment_AlreadyPaid(t *testing.T) {
	mockPaymentRepo := new(MockPaymentRepository)
	mockBookingRepo := new(MockBookingRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewPaymentService(mockPaymentRepo, mockBookingRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	booking := &domain.Booking{
		ID:     1,
		Status: "confirmed",
	}

	req := &domain.PaymentRequest{
		BookingID:     1,
		PaymentMethod: "CREDIT_CARD",
	}

	mockBookingRepo.On("GetByID", ctx, 1).Return(booking, nil)

	result, err := service.ProcessPayment(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "already paid")
	mockBookingRepo.AssertExpectations(t)
}

func TestPaymentService_ProcessPayment_BookingCancelled(t *testing.T) {
	mockPaymentRepo := new(MockPaymentRepository)
	mockBookingRepo := new(MockBookingRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewPaymentService(mockPaymentRepo, mockBookingRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	booking := &domain.Booking{
		ID:     1,
		Status: "cancelled",
	}

	req := &domain.PaymentRequest{
		BookingID:     1,
		PaymentMethod: "CREDIT_CARD",
	}

	mockBookingRepo.On("GetByID", ctx, 1).Return(booking, nil)

	result, err := service.ProcessPayment(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "cancelled")
	mockBookingRepo.AssertExpectations(t)
}

func TestPaymentService_ProcessPayment_InvalidPaymentMethod(t *testing.T) {
	mockPaymentRepo := new(MockPaymentRepository)
	mockBookingRepo := new(MockBookingRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewPaymentService(mockPaymentRepo, mockBookingRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	booking := &domain.Booking{
		ID:     1,
		Status: "pending",
	}

	req := &domain.PaymentRequest{
		BookingID:     1,
		PaymentMethod: "INVALID",
	}

	mockBookingRepo.On("GetByID", ctx, 1).Return(booking, nil)
	mockPaymentMethodRepo.On("GetByCode", ctx, "INVALID").Return(nil, errors.New("not found"))

	result, err := service.ProcessPayment(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid payment method")
	mockBookingRepo.AssertExpectations(t)
	mockPaymentMethodRepo.AssertExpectations(t)
}
