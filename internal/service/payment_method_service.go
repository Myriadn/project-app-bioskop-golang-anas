package service

import (
	"context"
	"fmt"

	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/repository"

	"go.uber.org/zap"
)

type PaymentMethodService interface {
	GetAllPaymentMethods(ctx context.Context) ([]*domain.PaymentMethod, error)
}

type paymentMethodService struct {
	paymentMethodRepo repository.PaymentMethodRepository
	logger            *zap.Logger
}

func NewPaymentMethodService(
	paymentMethodRepo repository.PaymentMethodRepository,
	logger *zap.Logger,
) PaymentMethodService {
	return &paymentMethodService{
		paymentMethodRepo: paymentMethodRepo,
		logger:            logger,
	}
}

func (s *paymentMethodService) GetAllPaymentMethods(ctx context.Context) ([]*domain.PaymentMethod, error) {
	methods, err := s.paymentMethodRepo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Failed to get payment methods", zap.Error(err))
		return nil, fmt.Errorf("failed to get payment methods: %w", err)
	}

	return methods, nil
}
