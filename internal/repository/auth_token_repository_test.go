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

func TestAuthTokenRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewAuthTokenRepository(mock)

	token := &domain.AuthToken{
		UserID:    1,
		Token:     "test-token-123",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "created_at"}).
		AddRow(1, now)

	mock.ExpectQuery("INSERT INTO auth_tokens").
		WithArgs(token.UserID, token.Token, token.ExpiresAt, pgxmock.AnyArg()).
		WillReturnRows(rows)

	err = repo.Create(context.Background(), token)

	assert.NoError(t, err)
	assert.Equal(t, 1, token.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthTokenRepository_GetByToken(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewAuthTokenRepository(mock)

	now := time.Now()
	expiresAt := now.Add(24 * time.Hour)

	rows := pgxmock.NewRows([]string{"id", "user_id", "token", "expires_at", "created_at"}).
		AddRow(1, 1, "test-token-123", expiresAt, now)

	mock.ExpectQuery("SELECT (.+) FROM auth_tokens WHERE token").
		WithArgs("test-token-123").
		WillReturnRows(rows)

	token, err := repo.GetByToken(context.Background(), "test-token-123")

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, "test-token-123", token.Token)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthTokenRepository_GetByToken_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewAuthTokenRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM auth_tokens WHERE token").
		WithArgs("invalid-token").
		WillReturnError(pgx.ErrNoRows)

	token, err := repo.GetByToken(context.Background(), "invalid-token")

	assert.Error(t, err)
	assert.Nil(t, token)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthTokenRepository_Delete(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewAuthTokenRepository(mock)

	mock.ExpectExec("DELETE FROM auth_tokens WHERE token").
		WithArgs("test-token-123").
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = repo.Delete(context.Background(), "test-token-123")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthTokenRepository_DeleteExpired(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewAuthTokenRepository(mock)

	mock.ExpectExec("DELETE FROM auth_tokens WHERE expires_at").
		WillReturnResult(pgxmock.NewResult("DELETE", 5))

	err = repo.DeleteExpired(context.Background())

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
