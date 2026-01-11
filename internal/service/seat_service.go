package service

import (
	"context"
	"fmt"

	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/repository"

	"go.uber.org/zap"
)

type SeatService interface {
	GetSeatAvailability(ctx context.Context, cinemaID int, date, time string) ([]*domain.SeatAvailability, *domain.Showtime, error)
}

type seatService struct {
	seatRepo     repository.SeatRepository
	showtimeRepo repository.ShowtimeRepository
	cinemaRepo   repository.CinemaRepository
	logger       *zap.Logger
}

func NewSeatService(
	seatRepo repository.SeatRepository,
	showtimeRepo repository.ShowtimeRepository,
	cinemaRepo repository.CinemaRepository,
	logger *zap.Logger,
) SeatService {
	return &seatService{
		seatRepo:     seatRepo,
		showtimeRepo: showtimeRepo,
		cinemaRepo:   cinemaRepo,
		logger:       logger,
	}
}

func (s *seatService) GetSeatAvailability(ctx context.Context, cinemaID int, date, time string) ([]*domain.SeatAvailability, *domain.Showtime, error) {
	// Validate cinema exists
	_, err := s.cinemaRepo.GetByID(ctx, cinemaID)
	if err != nil {
		return nil, nil, fmt.Errorf("cinema not found")
	}

	// Log untuk debugging
	s.logger.Info("Getting showtime",
		zap.Int("cinema_id", cinemaID),
		zap.String("date", date),
		zap.String("time", time),
	)

	// Get showtime
	showtime, err := s.showtimeRepo.GetByCinemaDateTime(ctx, cinemaID, date, time)
	if err != nil {
		s.logger.Error("Showtime not found",
			zap.Int("cinema_id", cinemaID),
			zap.String("date", date),
			zap.String("time", time),
			zap.Error(err),
		)
		return nil, nil, fmt.Errorf("no showtime found for the specified date and time")
	}

	s.logger.Info("Showtime found", zap.Int("showtime_id", showtime.ID))

	// Get seat availability
	seats, err := s.seatRepo.GetAvailableSeats(ctx, cinemaID, showtime.ID)
	if err != nil {
		s.logger.Error("Failed to get seat availability", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to get seat availability: %w", err)
	}

	return seats, showtime, nil
}
