package service

import (
	"context"
	"errors"
	"testing"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

type MockCinemaRepository struct {
	mock.Mock
}

func (m *MockCinemaRepository) GetAll(ctx context.Context, limit, offset int) ([]*domain.Cinema, int, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*domain.Cinema), args.Int(1), args.Error(2)
}

func (m *MockCinemaRepository) GetByID(ctx context.Context, id int) (*domain.Cinema, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Cinema), args.Error(1)
}

func TestCinemaService_GetAllCinemas(t *testing.T) {
	mockRepo := new(MockCinemaRepository)
	logger, _ := zap.NewDevelopment()
	service := NewCinemaService(mockRepo, logger)

	cinemas := []*domain.Cinema{
		{ID: 1, Name: "Cinema 1"},
		{ID: 2, Name: "Cinema 2"},
	}

	mockRepo.On("GetAll", mock.Anything, 10, 0).Return(cinemas, 2, nil)

	result, meta, err := service.GetAllCinemas(context.Background(), 1, 10)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, 2, meta.TotalRows)
	assert.Equal(t, 1, meta.Page)
	mockRepo.AssertExpectations(t)
}

func TestCinemaService_GetCinemaByID(t *testing.T) {
	mockRepo := new(MockCinemaRepository)
	logger, _ := zap.NewDevelopment()
	service := NewCinemaService(mockRepo, logger)

	cinema := &domain.Cinema{ID: 1, Name: "CGV Grand Indonesia"}
	mockRepo.On("GetByID", mock.Anything, 1).Return(cinema, nil)

	result, err := service.GetCinemaByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "CGV Grand Indonesia", result.Name)
	mockRepo.AssertExpectations(t)
}

// func TestCinemaService_GetAllCinemas_Error(t *testing.T) {
// 	mockRepo := new(MockCinemaRepository)
// 	logger := zap.NewNop()
// 	service := NewCinemaService(mockRepo, logger)

// 	mockRepo.On("GetAll", mock.Anything, 10, 0).Return(nil, 0, errors.New("database error"))

// 	result, meta, err := service.GetAllCinemas(context.Background(), 1, 10)

// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// 	assert.Nil(t, meta)
// 	mockRepo.AssertExpectations(t)
// }

func TestCinemaService_GetAllCinemas_EmptyResult(t *testing.T) {
	mockRepo := new(MockCinemaRepository)
	logger := zap.NewNop()
	service := NewCinemaService(mockRepo, logger)

	emptyCinemas := []*domain.Cinema{}

	mockRepo.On("GetAll", mock.Anything, 10, 0).Return(emptyCinemas, 0, nil)

	result, meta, err := service.GetAllCinemas(context.Background(), 1, 10)

	assert.NoError(t, err)
	assert.Len(t, result, 0)
	assert.Equal(t, 0, meta.TotalRows)
	mockRepo.AssertExpectations(t)
}

func TestCinemaService_GetCinemaByID_NotFound(t *testing.T) {
	mockRepo := new(MockCinemaRepository)
	logger := zap.NewNop()
	service := NewCinemaService(mockRepo, logger)

	mockRepo.On("GetByID", mock.Anything, 999).Return(nil, errors.New("cinema not found"))

	result, err := service.GetCinemaByID(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
