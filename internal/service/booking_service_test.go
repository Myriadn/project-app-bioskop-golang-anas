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

// Mock repositories
type MockBookingRepository struct {
	mock.Mock
}

func (m *MockBookingRepository) Create(ctx context.Context, booking *domain.Booking) error {
	args := m.Called(ctx, booking)
	return args.Error(0)
}

func (m *MockBookingRepository) GetByID(ctx context.Context, id int) (*domain.Booking, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Booking), args.Error(1)
}

func (m *MockBookingRepository) GetByUserID(ctx context.Context, userID int) ([]*domain.Booking, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Booking), args.Error(1)
}

func (m *MockBookingRepository) Update(ctx context.Context, booking *domain.Booking) error {
	args := m.Called(ctx, booking)
	return args.Error(0)
}

func (m *MockBookingRepository) CheckSeatBooked(ctx context.Context, showtimeID, seatID int) (bool, error) {
	args := m.Called(ctx, showtimeID, seatID)
	return args.Bool(0), args.Error(1)
}

type MockShowtimeRepository struct {
	mock.Mock
}

func (m *MockShowtimeRepository) GetByCinemaDateTime(ctx context.Context, cinemaID int, date, time string) (*domain.Showtime, error) {
	args := m.Called(ctx, cinemaID, date, time)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Showtime), args.Error(1)
}

func (m *MockShowtimeRepository) GetByID(ctx context.Context, id int) (*domain.Showtime, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Showtime), args.Error(1)
}

type MockSeatRepository struct {
	mock.Mock
}

func (m *MockSeatRepository) GetByCinemaID(ctx context.Context, cinemaID int) ([]*domain.Seat, error) {
	args := m.Called(ctx, cinemaID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Seat), args.Error(1)
}

func (m *MockSeatRepository) GetByID(ctx context.Context, id int) (*domain.Seat, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Seat), args.Error(1)
}

func (m *MockSeatRepository) GetAvailableSeats(ctx context.Context, cinemaID, showtimeID int) ([]*domain.SeatAvailability, error) {
	args := m.Called(ctx, cinemaID, showtimeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.SeatAvailability), args.Error(1)
}

type MockPaymentMethodRepository struct {
	mock.Mock
}

func (m *MockPaymentMethodRepository) GetAll(ctx context.Context) ([]*domain.PaymentMethod, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.PaymentMethod), args.Error(1)
}

func (m *MockPaymentMethodRepository) GetByCode(ctx context.Context, code string) (*domain.PaymentMethod, error) {
	args := m.Called(ctx, code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PaymentMethod), args.Error(1)
}

func (m *MockPaymentMethodRepository) GetByID(ctx context.Context, id int) (*domain.PaymentMethod, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PaymentMethod), args.Error(1)
}

func TestBookingService_CreateBooking_Success(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()
	now := time.Now()

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		MovieID:  1,
		Price:    50000,
	}

	seat := &domain.Seat{
		ID:       10,
		CinemaID: 1,
		SeatRow:  "A",
	}

	paymentMethod := &domain.PaymentMethod{
		ID:   1,
		Code: "CREDIT_CARD",
		Name: "Credit Card",
	}

	booking := &domain.Booking{
		ID:          1,
		UserID:      1,
		ShowtimeID:  1,
		SeatID:      10,
		BookingCode: "BK123456",
		Status:      "pending",
		TotalPrice:  50000,
		CreatedAt:   now,
	}

	req := &domain.BookingRequest{
		CinemaID:      1,
		SeatID:        10,
		Date:          "2024-01-15",
		Time:          "14:00",
		PaymentMethod: "CREDIT_CARD",
	}

	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetByID", ctx, 10).Return(seat, nil)
	mockBookingRepo.On("CheckSeatBooked", ctx, 1, 10).Return(false, nil)
	mockPaymentMethodRepo.On("GetByCode", ctx, "CREDIT_CARD").Return(paymentMethod, nil)
	mockBookingRepo.On("Create", ctx, mock.AnythingOfType("*domain.Booking")).Return(nil).Run(func(args mock.Arguments) {
		b := args.Get(1).(*domain.Booking)
		b.ID = 1
	})
	mockBookingRepo.On("GetByID", ctx, 1).Return(booking, nil)

	result, err := service.CreateBooking(ctx, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
	mockBookingRepo.AssertExpectations(t)
	mockPaymentMethodRepo.AssertExpectations(t)
}

func TestBookingService_CreateBooking_ShowtimeNotFound(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()
	req := &domain.BookingRequest{
		CinemaID:      1,
		SeatID:        10,
		Date:          "2024-01-15",
		Time:          "14:00",
		PaymentMethod: "CREDIT_CARD",
	}

	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(nil, errors.New("not found"))

	result, err := service.CreateBooking(ctx, 1, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "showtime not found")
	mockShowtimeRepo.AssertExpectations(t)
}

func TestBookingService_CreateBooking_SeatAlreadyBooked(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		Price:    50000,
	}

	seat := &domain.Seat{
		ID:       10,
		CinemaID: 1,
	}

	req := &domain.BookingRequest{
		CinemaID:      1,
		SeatID:        10,
		Date:          "2024-01-15",
		Time:          "14:00",
		PaymentMethod: "CREDIT_CARD",
	}

	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetByID", ctx, 10).Return(seat, nil)
	mockBookingRepo.On("CheckSeatBooked", ctx, 1, 10).Return(true, nil)

	result, err := service.CreateBooking(ctx, 1, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "seat is already booked")
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
	mockBookingRepo.AssertExpectations(t)
}

func TestBookingService_GetBookingByID_Success(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()
	booking := &domain.Booking{
		ID:          1,
		BookingCode: "BK123",
	}

	mockBookingRepo.On("GetByID", ctx, 1).Return(booking, nil)

	result, err := service.GetBookingByID(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "BK123", result.BookingCode)
	mockBookingRepo.AssertExpectations(t)
}

func TestBookingService_CreateBooking_SeatNotFound(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		Price:    50000,
	}

	req := &domain.BookingRequest{
		CinemaID:      1,
		SeatID:        999,
		Date:          "2024-01-15",
		Time:          "14:00",
		PaymentMethod: "CREDIT_CARD",
	}

	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetByID", ctx, 999).Return(nil, errors.New("not found"))

	result, err := service.CreateBooking(ctx, 1, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "seat not found")
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
}

func TestBookingService_CreateBooking_SeatNotInCinema(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		Price:    50000,
	}

	seat := &domain.Seat{
		ID:       10,
		CinemaID: 2, // Different cinema!
		SeatRow:  "A",
	}

	req := &domain.BookingRequest{
		CinemaID:      1,
		SeatID:        10,
		Date:          "2024-01-15",
		Time:          "14:00",
		PaymentMethod: "CREDIT_CARD",
	}

	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetByID", ctx, 10).Return(seat, nil)

	result, err := service.CreateBooking(ctx, 1, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "seat does not belong to this cinema")
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
}

func TestBookingService_CreateBooking_InvalidPaymentMethod(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		Price:    50000,
	}

	seat := &domain.Seat{
		ID:       10,
		CinemaID: 1,
		SeatRow:  "A",
	}

	req := &domain.BookingRequest{
		CinemaID:      1,
		SeatID:        10,
		Date:          "2024-01-15",
		Time:          "14:00",
		PaymentMethod: "INVALID_METHOD",
	}

	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetByID", ctx, 10).Return(seat, nil)
	mockBookingRepo.On("CheckSeatBooked", ctx, 1, 10).Return(false, nil)
	mockPaymentMethodRepo.On("GetByCode", ctx, "INVALID_METHOD").Return(nil, errors.New("not found"))

	result, err := service.CreateBooking(ctx, 1, req)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "invalid payment method")
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
	mockBookingRepo.AssertExpectations(t)
	mockPaymentMethodRepo.AssertExpectations(t)
}

func TestBookingService_GetBookingByID_NotFound(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	mockBookingRepo.On("GetByID", ctx, 999).Return(nil, errors.New("not found"))

	result, err := service.GetBookingByID(ctx, 999)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "booking not found")
	mockBookingRepo.AssertExpectations(t)
}

func TestBookingService_GetUserBookings_Success(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()
	bookings := []*domain.Booking{
		{ID: 1, BookingCode: "BK001"},
		{ID: 2, BookingCode: "BK002"},
	}

	mockBookingRepo.On("GetByUserID", ctx, 1).Return(bookings, nil)

	result, err := service.GetUserBookings(ctx, 1)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockBookingRepo.AssertExpectations(t)
}

func TestBookingService_GetUserBookings_Error(t *testing.T) {
	mockBookingRepo := new(MockBookingRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockSeatRepo := new(MockSeatRepository)
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewBookingService(mockBookingRepo, mockShowtimeRepo, mockSeatRepo, mockPaymentMethodRepo, logger)

	ctx := context.Background()

	mockBookingRepo.On("GetByUserID", ctx, 1).Return(nil, errors.New("database error"))

	result, err := service.GetUserBookings(ctx, 1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockBookingRepo.AssertExpectations(t)
}
