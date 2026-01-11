package repository

import (
	"context"
	"testing"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentRepository(mock)

	now := time.Now()
	payment := &domain.Payment{
		BookingID:       1,
		PaymentMethodID: 1,
		Amount:          50000,
		Status:          "completed",
		PaymentDetails:  domain.PaymentDetails{"card_type": "Visa"},
		PaidAt:          &now,
	}

	rows := pgxmock.NewRows([]string{"id", "created_at"}).
		AddRow(100, now)

	mock.ExpectQuery("INSERT INTO payments").
		WithArgs(
			payment.BookingID,
			payment.PaymentMethodID,
			payment.Amount,
			payment.Status,
			pgxmock.AnyArg(),
			payment.PaidAt,
			pgxmock.AnyArg(),
		).
		WillReturnRows(rows)

	err = repo.Create(context.Background(), payment)

	assert.NoError(t, err)
	assert.Equal(t, 100, payment.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPaymentRepository_GetByBookingID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{
		"id", "booking_id", "payment_method_id", "amount", "status", "payment_details", "paid_at", "created_at",
		"pm_id", "name", "code", "is_active", "pm_created_at",
	}).AddRow(
		100, 1, 1, 50000.0, "completed", []byte(`{"card_type":"Visa"}`), &now, now,
		1, "Credit Card", "CREDIT_CARD", true, now,
	)

	mock.ExpectQuery("SELECT (.+) FROM payments p JOIN payment_methods").
		WithArgs(1).
		WillReturnRows(rows)

	payment, err := repo.GetByBookingID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, 100, payment.ID)
	assert.Equal(t, "completed", payment.Status)
	assert.NotNil(t, payment.PaymentMethod)
	assert.Equal(t, "Credit Card", payment.PaymentMethod.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPaymentRepository_Update(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentRepository(mock)

	now := time.Now()
	payment := &domain.Payment{
		ID:     100,
		Status: "completed",
		PaidAt: &now,
	}

	mock.ExpectExec("UPDATE payments").
		WithArgs(payment.Status, payment.PaidAt, payment.ID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.Update(context.Background(), payment)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPaymentRepository_Create_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentRepository(mock)

	now := time.Now()
	payment := &domain.Payment{
		BookingID:       1,
		PaymentMethodID: 1,
		Amount:          50000,
		Status:          "completed",
		PaidAt:          &now,
	}

	mock.ExpectQuery("INSERT INTO payments").
		WithArgs(
			payment.BookingID,
			payment.PaymentMethodID,
			payment.Amount,
			payment.Status,
			pgxmock.AnyArg(),
			payment.PaidAt,
			pgxmock.AnyArg(),
		).
		WillReturnError(assert.AnError)

	err = repo.Create(context.Background(), payment)

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
