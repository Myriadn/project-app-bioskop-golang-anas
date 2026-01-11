package service

import (
	"context"
	"errors"
	"testing"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestSeatService_GetSeatAvailability_Success(t *testing.T) {
	mockSeatRepo := new(MockSeatRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockCinemaRepo := new(MockCinemaRepository)
	logger := zap.NewNop()

	service := NewSeatService(mockSeatRepo, mockShowtimeRepo, mockCinemaRepo, logger)

	ctx := context.Background()

	cinema := &domain.Cinema{
		ID:   1,
		Name: "CGV",
	}

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		MovieID:  1,
	}

	seats := []*domain.SeatAvailability{
		{
			Seat: &domain.Seat{
				ID:       1,
				CinemaID: 1,
				SeatRow:  "A",
			},
			IsBooked:   false,
			ShowtimeID: 1,
		},
		{
			Seat: &domain.Seat{
				ID:       2,
				CinemaID: 1,
				SeatRow:  "A",
			},
			IsBooked:   true,
			ShowtimeID: 1,
		},
	}

	mockCinemaRepo.On("GetByID", ctx, 1).Return(cinema, nil)
	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetAvailableSeats", ctx, 1, 1).Return(seats, nil)

	resultSeats, resultShowtime, err := service.GetSeatAvailability(ctx, 1, "2024-01-15", "14:00")

	assert.NoError(t, err)
	assert.NotNil(t, resultSeats)
	assert.NotNil(t, resultShowtime)
	assert.Len(t, resultSeats, 2)
	assert.False(t, resultSeats[0].IsBooked)
	assert.True(t, resultSeats[1].IsBooked)
	mockCinemaRepo.AssertExpectations(t)
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
}

func TestSeatService_GetSeatAvailability_CinemaNotFound(t *testing.T) {
	mockSeatRepo := new(MockSeatRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockCinemaRepo := new(MockCinemaRepository)
	logger := zap.NewNop()

	service := NewSeatService(mockSeatRepo, mockShowtimeRepo, mockCinemaRepo, logger)

	ctx := context.Background()

	mockCinemaRepo.On("GetByID", ctx, 999).Return(nil, errors.New("not found"))

	resultSeats, resultShowtime, err := service.GetSeatAvailability(ctx, 999, "2024-01-15", "14:00")

	assert.Error(t, err)
	assert.Nil(t, resultSeats)
	assert.Nil(t, resultShowtime)
	assert.Contains(t, err.Error(), "cinema not found")
	mockCinemaRepo.AssertExpectations(t)
}

func TestSeatService_GetSeatAvailability_ShowtimeNotFound(t *testing.T) {
	mockSeatRepo := new(MockSeatRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockCinemaRepo := new(MockCinemaRepository)
	logger := zap.NewNop()

	service := NewSeatService(mockSeatRepo, mockShowtimeRepo, mockCinemaRepo, logger)

	ctx := context.Background()

	cinema := &domain.Cinema{
		ID:   1,
		Name: "CGV",
	}

	mockCinemaRepo.On("GetByID", ctx, 1).Return(cinema, nil)
	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(nil, errors.New("not found"))

	resultSeats, resultShowtime, err := service.GetSeatAvailability(ctx, 1, "2024-01-15", "14:00")

	assert.Error(t, err)
	assert.Nil(t, resultSeats)
	assert.Nil(t, resultShowtime)
	assert.Contains(t, err.Error(), "no showtime found")
	mockCinemaRepo.AssertExpectations(t)
	mockShowtimeRepo.AssertExpectations(t)
}

func TestSeatService_GetSeatAvailability_EmptySeats(t *testing.T) {
	mockSeatRepo := new(MockSeatRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockCinemaRepo := new(MockCinemaRepository)
	logger := zap.NewNop()

	service := NewSeatService(mockSeatRepo, mockShowtimeRepo, mockCinemaRepo, logger)

	ctx := context.Background()

	cinema := &domain.Cinema{
		ID:   1,
		Name: "CGV",
	}

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		MovieID:  1,
	}

	emptySeats := []*domain.SeatAvailability{}

	mockCinemaRepo.On("GetByID", ctx, 1).Return(cinema, nil)
	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetAvailableSeats", ctx, 1, 1).Return(emptySeats, nil)

	resultSeats, resultShowtime, err := service.GetSeatAvailability(ctx, 1, "2024-01-15", "14:00")

	assert.NoError(t, err)
	assert.NotNil(t, resultSeats)
	assert.NotNil(t, resultShowtime)
	assert.Len(t, resultSeats, 0)
	mockCinemaRepo.AssertExpectations(t)
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
}

func TestSeatService_GetSeatAvailability_GetSeatsError(t *testing.T) {
	mockSeatRepo := new(MockSeatRepository)
	mockShowtimeRepo := new(MockShowtimeRepository)
	mockCinemaRepo := new(MockCinemaRepository)
	logger := zap.NewNop()

	service := NewSeatService(mockSeatRepo, mockShowtimeRepo, mockCinemaRepo, logger)

	ctx := context.Background()

	cinema := &domain.Cinema{
		ID:   1,
		Name: "CGV",
	}

	showtime := &domain.Showtime{
		ID:       1,
		CinemaID: 1,
		MovieID:  1,
	}

	mockCinemaRepo.On("GetByID", ctx, 1).Return(cinema, nil)
	mockShowtimeRepo.On("GetByCinemaDateTime", ctx, 1, "2024-01-15", "14:00").Return(showtime, nil)
	mockSeatRepo.On("GetAvailableSeats", ctx, 1, 1).Return(nil, errors.New("database error"))

	resultSeats, resultShowtime, err := service.GetSeatAvailability(ctx, 1, "2024-01-15", "14:00")

	assert.Error(t, err)
	assert.Nil(t, resultSeats)
	assert.Nil(t, resultShowtime)
	assert.Contains(t, err.Error(), "failed to get seat availability")
	mockCinemaRepo.AssertExpectations(t)
	mockShowtimeRepo.AssertExpectations(t)
	mockSeatRepo.AssertExpectations(t)
}
