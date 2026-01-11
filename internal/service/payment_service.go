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

type PaymentService interface {
	ProcessPayment(ctx context.Context, req *domain.PaymentRequest) (*domain.Payment, error)
}

type paymentService struct {
	paymentRepo       repository.PaymentRepository
	bookingRepo       repository.BookingRepository
	paymentMethodRepo repository.PaymentMethodRepository
	logger            *zap.Logger
}

func NewPaymentService(
	paymentRepo repository.PaymentRepository,
	bookingRepo repository.BookingRepository,
	paymentMethodRepo repository.PaymentMethodRepository,
	logger *zap.Logger,
) PaymentService {
	return &paymentService{
		paymentRepo:       paymentRepo,
		bookingRepo:       bookingRepo,
		paymentMethodRepo: paymentMethodRepo,
		logger:            logger,
	}
}

func (s *paymentService) ProcessPayment(ctx context.Context, req *domain.PaymentRequest) (*domain.Payment, error) {
	// Validate booking exists
	booking, err := s.bookingRepo.GetByID(ctx, req.BookingID)
	if err != nil {
		s.logger.Error("Booking not found", zap.Error(err))
		return nil, fmt.Errorf("booking not found")
	}

	// Check if booking is already confirmed or cancelled
	if booking.Status == "confirmed" {
		return nil, fmt.Errorf("booking is already paid")
	}

	if booking.Status == "cancelled" {
		return nil, fmt.Errorf("booking is cancelled")
	}

	// Validate payment method
	paymentMethod, err := s.paymentMethodRepo.GetByCode(ctx, req.PaymentMethod)
	if err != nil {
		s.logger.Error("Invalid payment method", zap.Error(err))
		return nil, fmt.Errorf("invalid payment method")
	}

	// Validate payment details
	if req.PaymentDetails == nil {
		req.PaymentDetails = domain.PaymentDetails{}
	}

	// Add metadata if not provided
	if req.PaymentDetails["timestamp"] == nil {
		req.PaymentDetails["timestamp"] = time.Now().Format(time.RFC3339)
	}
	if req.PaymentDetails["payment_method"] == nil {
		req.PaymentDetails["payment_method"] = paymentMethod.Name
	}

	// Create payment record
	payment := &domain.Payment{
		BookingID:       req.BookingID,
		PaymentMethodID: paymentMethod.ID,
		Amount:          booking.TotalPrice,
		Status:          "success",
		PaymentDetails:  req.PaymentDetails,
	}

	now := time.Now()
	payment.PaidAt = &now

	if err := s.paymentRepo.Create(ctx, payment); err != nil {
		s.logger.Error("Failed to create payment", zap.Error(err))
		return nil, fmt.Errorf("failed to process payment: %w", err)
	}

	// Update booking status to confirmed
	booking.Status = "confirmed"
	if err := s.bookingRepo.Update(ctx, booking); err != nil {
		s.logger.Error("Failed to update booking status", zap.Error(err))
	}

	s.logger.Info("Payment processed successfully",
		zap.Int("payment_id", payment.ID),
		zap.Int("booking_id", req.BookingID),
		zap.Float64("amount", payment.Amount),
		zap.Any("payment_details", payment.PaymentDetails),
	)

	// GOROUTINE: Async logging to file
	utils.LogPaymentAsync(s.logger, req.BookingID, payment.Amount, paymentMethod.Name)

	// GOROUTINE: Async update booking activity log
	utils.LogBookingAsync(s.logger, booking.UserID, booking.BookingCode, "confirmed")

	// Get full payment details
	fullPayment, err := s.paymentRepo.GetByBookingID(ctx, req.BookingID)
	if err != nil {
		return payment, nil
	}

	return fullPayment, nil
}
