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

func TestUserRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewUserRepository(mock)

	user := &domain.User{
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "hashedpassword",
		IsVerified:   false,
	}

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(1, now, now)

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(user.Username, user.Email, user.PasswordHash, user.IsVerified, pgxmock.AnyArg(), pgxmock.AnyArg()).
		WillReturnRows(rows)

	err = repo.Create(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewUserRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "username", "email", "password_hash", "is_verified", "created_at", "updated_at"}).
		AddRow(1, "testuser", "test@example.com", "hashedpassword", true, now, now)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE id").
		WithArgs(1).
		WillReturnRows(rows)

	user, err := repo.GetByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByUsername(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewUserRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "username", "email", "password_hash", "is_verified", "created_at", "updated_at"}).
		AddRow(1, "testuser", "test@example.com", "hashedpassword", false, now, now)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE username").
		WithArgs("testuser").
		WillReturnRows(rows)

	user, err := repo.GetByUsername(context.Background(), "testuser")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByUsername_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewUserRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE username").
		WithArgs("nonexistent").
		WillReturnError(pgx.ErrNoRows)

	user, err := repo.GetByUsername(context.Background(), "nonexistent")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "user not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByEmail(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewUserRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "username", "email", "password_hash", "is_verified", "created_at", "updated_at"}).
		AddRow(1, "testuser", "test@example.com", "hashedpassword", false, now, now)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE email").
		WithArgs("test@example.com").
		WillReturnRows(rows)

	user, err := repo.GetByEmail(context.Background(), "test@example.com")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@example.com", user.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_GetByEmail_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewUserRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM users WHERE email").
		WithArgs("nonexistent@example.com").
		WillReturnError(pgx.ErrNoRows)

	user, err := repo.GetByEmail(context.Background(), "nonexistent@example.com")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "user not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserRepository_Update(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewUserRepository(mock)

	user := &domain.User{
		ID:         1,
		Username:   "updateduser",
		Email:      "updated@example.com",
		IsVerified: true,
	}

	mock.ExpectExec("UPDATE users").
		WithArgs(user.Username, user.Email, user.IsVerified, pgxmock.AnyArg(), user.ID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.Update(context.Background(), user)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
