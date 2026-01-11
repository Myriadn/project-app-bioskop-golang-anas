package repository

import (
	"context"
	"testing"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOTPRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewOTPRepository(mock)

	otp := &domain.OTPCode{
		UserID:    1,
		Code:      "123456",
		IsUsed:    false,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "created_at"}).
		AddRow(1, now)

	mock.ExpectQuery("INSERT INTO otp_codes").
		WithArgs(otp.UserID, otp.Code, otp.IsUsed, pgxmock.AnyArg(), pgxmock.AnyArg()).
		WillReturnRows(rows)

	err = repo.Create(context.Background(), otp)

	assert.NoError(t, err)
	assert.Equal(t, 1, otp.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOTPRepository_GetByUserIDAndCode(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewOTPRepository(mock)

	now := time.Now()
	expiresAt := now.Add(10 * time.Minute)

	rows := pgxmock.NewRows([]string{"id", "user_id", "code", "is_used", "expires_at", "created_at"}).
		AddRow(1, 1, "123456", false, expiresAt, now)

	mock.ExpectQuery("SELECT (.+) FROM otp_codes WHERE user_id").
		WithArgs(1, "123456").
		WillReturnRows(rows)

	otp, err := repo.GetByUserIDAndCode(context.Background(), 1, "123456")

	assert.NoError(t, err)
	assert.NotNil(t, otp)
	assert.Equal(t, "123456", otp.Code)
	assert.False(t, otp.IsUsed)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOTPRepository_GetByUserIDAndCode_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewOTPRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM otp_codes WHERE user_id").
		WithArgs(1, "999999").
		WillReturnError(pgx.ErrNoRows)

	otp, err := repo.GetByUserIDAndCode(context.Background(), 1, "999999")

	assert.Error(t, err)
	assert.Nil(t, otp)
	assert.Contains(t, err.Error(), "invalid or expired OTP")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOTPRepository_MarkAsUsed(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewOTPRepository(mock)

	mock.ExpectExec("UPDATE otp_codes SET is_used").
		WithArgs(1).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.MarkAsUsed(context.Background(), 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOTPRepository_DeleteExpired(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewOTPRepository(mock)

	mock.ExpectExec("DELETE FROM otp_codes WHERE expires_at").
		WillReturnResult(pgxmock.NewResult("DELETE", 5))

	err = repo.DeleteExpired(context.Background())

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOTPRepository_DeleteByUserID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewOTPRepository(mock)

	mock.ExpectExec("DELETE FROM otp_codes WHERE user_id").
		WithArgs(1).
		WillReturnResult(pgxmock.NewResult("DELETE", 2))

	err = repo.DeleteByUserID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
