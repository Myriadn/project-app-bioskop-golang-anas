package repository

import (
	"context"
	"fmt"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) error
	GetByBookingID(ctx context.Context, bookingID int) (*domain.Payment, error)
	Update(ctx context.Context, payment *domain.Payment) error
}

type paymentRepository struct {
	db *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	query := `
		INSERT INTO payments (booking_id, payment_method_id, amount, status, payment_details, paid_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	now := time.Now()
	err := r.db.QueryRow(
		ctx,
		query,
		payment.BookingID,
		payment.PaymentMethodID,
		payment.Amount,
		payment.Status,
		payment.PaymentDetails,
		payment.PaidAt,
		now,
	).Scan(&payment.ID, &payment.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}

	return nil
}

func (r *paymentRepository) GetByBookingID(ctx context.Context, bookingID int) (*domain.Payment, error) {
	query := `
		SELECT p.id, p.booking_id, p.payment_method_id, p.amount, p.status, p.payment_details, p.paid_at, p.created_at,
		       pm.id, pm.name, pm.code, pm.is_active, pm.created_at
		FROM payments p
		JOIN payment_methods pm ON p.payment_method_id = pm.id
		WHERE p.booking_id = $1
	`

	var payment domain.Payment
	var paymentMethod domain.PaymentMethod

	err := r.db.QueryRow(ctx, query, bookingID).Scan(
		&payment.ID,
		&payment.BookingID,
		&payment.PaymentMethodID,
		&payment.Amount,
		&payment.Status,
		&payment.PaymentDetails,
		&payment.PaidAt,
		&payment.CreatedAt,
		&paymentMethod.ID,
		&paymentMethod.Name,
		&paymentMethod.Code,
		&paymentMethod.IsActive,
		&paymentMethod.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("payment not found")
	}

	payment.PaymentMethod = &paymentMethod

	return &payment, nil
}

func (r *paymentRepository) Update(ctx context.Context, payment *domain.Payment) error {
	query := `
		UPDATE payments
		SET status = $1, paid_at = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(ctx, query, payment.Status, payment.PaidAt, payment.ID)
	if err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	return nil
}
