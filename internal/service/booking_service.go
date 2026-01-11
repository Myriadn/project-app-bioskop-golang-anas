package service

import (
	"context"
	"fmt"

	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/repository"
	"project-app-bioskop-golang-homework-anas/internal/utils"

	"go.uber.org/zap"
)

type BookingService interface {
	CreateBooking(ctx context.Context, userID int, req *domain.BookingRequest) (*domain.Booking, error)
	GetUserBookings(ctx context.Context, userID int) ([]*domain.Booking, error)
	GetBookingByID(ctx context.Context, bookingID int) (*domain.Booking, error)
}

type bookingService struct {
	bookingRepo       repository.BookingRepository
	showtimeRepo      repository.ShowtimeRepository
	seatRepo          repository.SeatRepository
	paymentMethodRepo repository.PaymentMethodRepository
	logger            *zap.Logger
}

func NewBookingService(
	bookingRepo repository.BookingRepository,
	showtimeRepo repository.ShowtimeRepository,
	seatRepo repository.SeatRepository,
	paymentMethodRepo repository.PaymentMethodRepository,
	logger *zap.Logger,
) BookingService {
	return &bookingService{
		bookingRepo:       bookingRepo,
		showtimeRepo:      showtimeRepo,
		seatRepo:          seatRepo,
		paymentMethodRepo: paymentMethodRepo,
		logger:            logger,
	}
}

func (s *bookingService) CreateBooking(ctx context.Context, userID int, req *domain.BookingRequest) (*domain.Booking, error) {
	// Validate showtime exists
	showtime, err := s.showtimeRepo.GetByCinemaDateTime(ctx, req.CinemaID, req.Date, req.Time)
	if err != nil {
		s.logger.Error("Showtime not found", zap.Error(err))
		return nil, fmt.Errorf("showtime not found for the specified date and time")
	}

	// Validate seat exists and belongs to cinema
	seat, err := s.seatRepo.GetByID(ctx, req.SeatID)
	if err != nil {
		s.logger.Error("Seat not found", zap.Error(err))
		return nil, fmt.Errorf("seat not found")
	}

	if seat.CinemaID != req.CinemaID {
		return nil, fmt.Errorf("seat does not belong to this cinema")
	}

	// Check if seat is already booked for this showtime
	isBooked, err := s.bookingRepo.CheckSeatBooked(ctx, showtime.ID, req.SeatID)
	if err != nil {
		s.logger.Error("Failed to check seat availability", zap.Error(err))
		return nil, fmt.Errorf("failed to check seat availability")
	}

	if isBooked {
		return nil, fmt.Errorf("seat is already booked for this showtime")
	}

	// Validate payment method
	_, err = s.paymentMethodRepo.GetByCode(ctx, req.PaymentMethod)
	if err != nil {
		s.logger.Error("Invalid payment method", zap.Error(err))
		return nil, fmt.Errorf("invalid payment method")
	}

	// Generate booking code
	bookingCode, err := utils.GenerateBookingCode()
	if err != nil {
		s.logger.Error("Failed to generate booking code", zap.Error(err))
		return nil, fmt.Errorf("failed to generate booking code")
	}

	// Create booking
	booking := &domain.Booking{
		UserID:      userID,
		ShowtimeID:  showtime.ID,
		SeatID:      req.SeatID,
		BookingCode: bookingCode,
		Status:      "pending",
		TotalPrice:  showtime.Price,
	}

	if err := s.bookingRepo.Create(ctx, booking); err != nil {
		s.logger.Error("Failed to create booking", zap.Error(err))
		return nil, fmt.Errorf("failed to create booking: %w", err)
	}

	s.logger.Info("Booking created successfully",
		zap.Int("booking_id", booking.ID),
		zap.Int("user_id", userID),
		zap.String("booking_code", bookingCode),
	)

	// Get full booking details
	fullBooking, err := s.bookingRepo.GetByID(ctx, booking.ID)
	if err != nil {
		s.logger.Error("Failed to get booking details", zap.Error(err))
		return booking, nil // Return basic booking if full details fail
	}

	return fullBooking, nil
}

func (s *bookingService) GetUserBookings(ctx context.Context, userID int) ([]*domain.Booking, error) {
	bookings, err := s.bookingRepo.GetByUserID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to get user bookings", zap.Int("user_id", userID), zap.Error(err))
		return nil, fmt.Errorf("failed to get bookings: %w", err)
	}

	return bookings, nil
}

func (s *bookingService) GetBookingByID(ctx context.Context, bookingID int) (*domain.Booking, error) {
	booking, err := s.bookingRepo.GetByID(ctx, bookingID)
	if err != nil {
		s.logger.Error("Failed to get booking", zap.Int("booking_id", bookingID), zap.Error(err))
		return nil, fmt.Errorf("booking not found")
	}

	return booking, nil
}
