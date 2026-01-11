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

func TestCinemaRepository_GetAll(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewCinemaRepository(mock)

	countRows := pgxmock.NewRows([]string{"count"}).AddRow(5)
	mock.ExpectQuery("SELECT COUNT").WillReturnRows(countRows)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "name", "location", "description", "created_at"}).
		AddRow(1, "CGV Grand Indonesia", "Jakarta Pusat", "Premium cinema", now).
		AddRow(2, "XXI Plaza Senayan", "Jakarta Selatan", "Modern cinema", now)

	mock.ExpectQuery("SELECT (.+) FROM cinemas").
		WithArgs(10, 0).
		WillReturnRows(rows)

	cinemas, total, err := repo.GetAll(context.Background(), 10, 0)

	assert.NoError(t, err)
	assert.Equal(t, 5, total)
	assert.Len(t, cinemas, 2)
	assert.Equal(t, "CGV Grand Indonesia", cinemas[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCinemaRepository_GetAll_Empty(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewCinemaRepository(mock)

	countRows := pgxmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT COUNT").WillReturnRows(countRows)

	rows := pgxmock.NewRows([]string{"id", "name", "location", "description", "created_at"})
	mock.ExpectQuery("SELECT (.+) FROM cinemas").
		WithArgs(10, 0).
		WillReturnRows(rows)

	cinemas, total, err := repo.GetAll(context.Background(), 10, 0)

	assert.NoError(t, err)
	assert.Equal(t, 0, total)
	assert.Len(t, cinemas, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCinemaRepository_GetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewCinemaRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "name", "location", "description", "created_at"}).
		AddRow(1, "CGV Grand Indonesia", "Jakarta Pusat", "Premium cinema", now)

	mock.ExpectQuery("SELECT (.+) FROM cinemas WHERE id").
		WithArgs(1).
		WillReturnRows(rows)

	cinema, err := repo.GetByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, cinema)
	assert.Equal(t, "CGV Grand Indonesia", cinema.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCinemaRepository_GetByID_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewCinemaRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM cinemas WHERE id").
		WithArgs(999).
		WillReturnError(pgx.ErrNoRows)

	cinema, err := repo.GetByID(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, cinema)
	assert.Contains(t, err.Error(), "cinema not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}
