package repository

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentMethodRepository_GetAll(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentMethodRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "name", "code", "is_active", "created_at"}).
		AddRow(1, "Credit Card", "CREDIT_CARD", true, now).
		AddRow(2, "GoPay", "GOPAY", true, now).
		AddRow(3, "Bank Transfer", "BANK_TRANSFER", true, now)

	mock.ExpectQuery("SELECT (.+) FROM payment_methods WHERE is_active").
		WillReturnRows(rows)

	methods, err := repo.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, methods, 3)
	assert.Equal(t, "Credit Card", methods[0].Name)
	assert.Equal(t, "CREDIT_CARD", methods[0].Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPaymentMethodRepository_GetAll_Empty(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentMethodRepository(mock)

	rows := pgxmock.NewRows([]string{"id", "name", "code", "is_active", "created_at"})

	mock.ExpectQuery("SELECT (.+) FROM payment_methods WHERE is_active").
		WillReturnRows(rows)

	methods, err := repo.GetAll(context.Background())

	assert.NoError(t, err)
	assert.Len(t, methods, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPaymentMethodRepository_GetByCode(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentMethodRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "name", "code", "is_active", "created_at"}).
		AddRow(1, "Credit Card", "CREDIT_CARD", true, now)

	mock.ExpectQuery("SELECT (.+) FROM payment_methods WHERE code").
		WithArgs("CREDIT_CARD").
		WillReturnRows(rows)

	method, err := repo.GetByCode(context.Background(), "CREDIT_CARD")

	assert.NoError(t, err)
	assert.NotNil(t, method)
	assert.Equal(t, "Credit Card", method.Name)
	assert.Equal(t, "CREDIT_CARD", method.Code)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPaymentMethodRepository_GetByCode_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentMethodRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM payment_methods WHERE code").
		WithArgs("INVALID").
		WillReturnError(pgx.ErrNoRows)

	method, err := repo.GetByCode(context.Background(), "INVALID")

	assert.Error(t, err)
	assert.Nil(t, method)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPaymentMethodRepository_GetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewPaymentMethodRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "name", "code", "is_active", "created_at"}).
		AddRow(1, "Credit Card", "CREDIT_CARD", true, now)

	mock.ExpectQuery("SELECT (.+) FROM payment_methods WHERE id").
		WithArgs(1).
		WillReturnRows(rows)

	method, err := repo.GetByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, method)
	assert.Equal(t, 1, method.ID)
	assert.Equal(t, "Credit Card", method.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
