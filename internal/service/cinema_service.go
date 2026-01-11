package service

import (
	"context"
	"fmt"
	"math"

	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/repository"
	"project-app-bioskop-golang-homework-anas/internal/utils"

	"go.uber.org/zap"
)

type CinemaService interface {
	GetAllCinemas(ctx context.Context, page, limit int) ([]*domain.Cinema, *utils.PaginationMeta, error)
	GetCinemaByID(ctx context.Context, id int) (*domain.Cinema, error)
}

type cinemaService struct {
	cinemaRepo repository.CinemaRepository
	logger     *zap.Logger
}

func NewCinemaService(cinemaRepo repository.CinemaRepository, logger *zap.Logger) CinemaService {
	return &cinemaService{
		cinemaRepo: cinemaRepo,
		logger:     logger,
	}
}

func (s *cinemaService) GetAllCinemas(ctx context.Context, page, limit int) ([]*domain.Cinema, *utils.PaginationMeta, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10 // default limit
	}

	offset := (page - 1) * limit

	// Get cinemas
	cinemas, total, err := s.cinemaRepo.GetAll(ctx, limit, offset)
	if err != nil {
		s.logger.Error("Failed to get cinemas", zap.Error(err))
		return nil, nil, fmt.Errorf("failed to get cinemas: %w", err)
	}

	// Calculate pagination meta
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	meta := &utils.PaginationMeta{
		Page:       page,
		Limit:      limit,
		TotalRows:  total,
		TotalPages: totalPages,
	}

	return cinemas, meta, nil
}

func (s *cinemaService) GetCinemaByID(ctx context.Context, id int) (*domain.Cinema, error) {
	cinema, err := s.cinemaRepo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get cinema", zap.Int("cinema_id", id), zap.Error(err))
		return nil, fmt.Errorf("cinema not found")
	}

	return cinema, nil
}
