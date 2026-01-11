package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPaymentMethodService_GetAllPaymentMethods_Success(t *testing.T) {
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewPaymentMethodService(mockPaymentMethodRepo, logger)

	ctx := context.Background()
	now := time.Now()

	methods := []*domain.PaymentMethod{
		{
			ID:        1,
			Name:      "Credit Card",
			Code:      "CREDIT_CARD",
			IsActive:  true,
			CreatedAt: now,
		},
		{
			ID:        2,
			Name:      "GoPay",
			Code:      "GOPAY",
			IsActive:  true,
			CreatedAt: now,
		},
	}

	mockPaymentMethodRepo.On("GetAll", ctx).Return(methods, nil)

	result, err := service.GetAllPaymentMethods(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Credit Card", result[0].Name)
	mockPaymentMethodRepo.AssertExpectations(t)
}

func TestPaymentMethodService_GetAllPaymentMethods_Error(t *testing.T) {
	mockPaymentMethodRepo := new(MockPaymentMethodRepository)
	logger := zap.NewNop()

	service := NewPaymentMethodService(mockPaymentMethodRepo, logger)

	ctx := context.Background()

	mockPaymentMethodRepo.On("GetAll", ctx).Return(nil, errors.New("database error"))

	result, err := service.GetAllPaymentMethods(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockPaymentMethodRepo.AssertExpectations(t)
}
