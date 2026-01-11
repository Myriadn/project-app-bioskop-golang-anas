package repository

import (
	"context"
	"fmt"

	"project-app-bioskop-golang-homework-anas/internal/domain"
)

type PaymentMethodRepository interface {
	GetAll(ctx context.Context) ([]*domain.PaymentMethod, error)
	GetByCode(ctx context.Context, code string) (*domain.PaymentMethod, error)
	GetByID(ctx context.Context, id int) (*domain.PaymentMethod, error)
}

type paymentMethodRepository struct {
	db PgxPool
}

func NewPaymentMethodRepository(db PgxPool) PaymentMethodRepository {
	return &paymentMethodRepository{db: db}
}

func (r *paymentMethodRepository) GetAll(ctx context.Context) ([]*domain.PaymentMethod, error) {
	query := `
		SELECT id, name, code, is_active, created_at
		FROM payment_methods
		WHERE is_active = true
		ORDER BY name ASC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment methods: %w", err)
	}
	defer rows.Close()

	var methods []*domain.PaymentMethod
	for rows.Next() {
		var method domain.PaymentMethod
		err := rows.Scan(
			&method.ID,
			&method.Name,
			&method.Code,
			&method.IsActive,
			&method.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan payment method: %w", err)
		}
		methods = append(methods, &method)
	}

	return methods, nil
}

func (r *paymentMethodRepository) GetByCode(ctx context.Context, code string) (*domain.PaymentMethod, error) {
	query := `
		SELECT id, name, code, is_active, created_at
		FROM payment_methods
		WHERE code = $1 AND is_active = true
	`

	var method domain.PaymentMethod
	err := r.db.QueryRow(ctx, query, code).Scan(
		&method.ID,
		&method.Name,
		&method.Code,
		&method.IsActive,
		&method.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("payment method not found")
	}

	return &method, nil
}

func (r *paymentMethodRepository) GetByID(ctx context.Context, id int) (*domain.PaymentMethod, error) {
	query := `
		SELECT id, name, code, is_active, created_at
		FROM payment_methods
		WHERE id = $1
	`

	var method domain.PaymentMethod
	err := r.db.QueryRow(ctx, query, id).Scan(
		&method.ID,
		&method.Name,
		&method.Code,
		&method.IsActive,
		&method.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("payment method not found")
	}

	return &method, nil
}
