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

func TestSeatRepository_GetByCinemaID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "cinema_id", "seat_row", "seat_number", "seat_type", "created_at"}).
		AddRow(1, 1, "A", 1, "regular", now).
		AddRow(2, 1, "A", 2, "regular", now).
		AddRow(3, 1, "B", 1, "vip", now)

	mock.ExpectQuery("SELECT (.+) FROM seats WHERE cinema_id").
		WithArgs(1).
		WillReturnRows(rows)

	seats, err := repo.GetByCinemaID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, seats, 3)
	assert.Equal(t, "A", seats[0].SeatRow)
	assert.Equal(t, 1, seats[0].SeatNumber)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeatRepository_GetByCinemaID_Empty(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	rows := pgxmock.NewRows([]string{"id", "cinema_id", "seat_row", "seat_number", "seat_type", "created_at"})

	mock.ExpectQuery("SELECT (.+) FROM seats WHERE cinema_id").
		WithArgs(999).
		WillReturnRows(rows)

	seats, err := repo.GetByCinemaID(context.Background(), 999)

	assert.NoError(t, err)
	assert.Len(t, seats, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeatRepository_GetByCinemaID_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM seats WHERE cinema_id").
		WithArgs(1).
		WillReturnError(assert.AnError)

	seats, err := repo.GetByCinemaID(context.Background(), 1)

	assert.Error(t, err)
	assert.Nil(t, seats)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeatRepository_GetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "cinema_id", "seat_row", "seat_number", "seat_type", "created_at"}).
		AddRow(10, 1, "C", 5, "vip", now)

	mock.ExpectQuery("SELECT (.+) FROM seats WHERE id").
		WithArgs(10).
		WillReturnRows(rows)

	seat, err := repo.GetByID(context.Background(), 10)

	assert.NoError(t, err)
	assert.NotNil(t, seat)
	assert.Equal(t, 10, seat.ID)
	assert.Equal(t, "C", seat.SeatRow)
	assert.Equal(t, 5, seat.SeatNumber)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeatRepository_GetByID_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM seats WHERE id").
		WithArgs(999).
		WillReturnError(pgx.ErrNoRows)

	seat, err := repo.GetByID(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, seat)
	assert.Contains(t, err.Error(), "seat not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeatRepository_GetAvailableSeats(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "cinema_id", "seat_row", "seat_number", "seat_type", "created_at", "is_booked"}).
		AddRow(1, 1, "A", 1, "regular", now, false).
		AddRow(2, 1, "A", 2, "regular", now, true).
		AddRow(3, 1, "B", 1, "vip", now, false)

	mock.ExpectQuery("SELECT (.+) FROM seats s LEFT JOIN bookings").
		WithArgs(1, 10).
		WillReturnRows(rows)

	availableSeats, err := repo.GetAvailableSeats(context.Background(), 1, 10)

	assert.NoError(t, err)
	assert.Len(t, availableSeats, 3)
	assert.False(t, availableSeats[0].IsBooked)
	assert.True(t, availableSeats[1].IsBooked)
	assert.Equal(t, 10, availableSeats[0].ShowtimeID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeatRepository_GetAvailableSeats_AllBooked(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "cinema_id", "seat_row", "seat_number", "seat_type", "created_at", "is_booked"}).
		AddRow(1, 1, "A", 1, "regular", now, true).
		AddRow(2, 1, "A", 2, "regular", now, true)

	mock.ExpectQuery("SELECT (.+) FROM seats s LEFT JOIN bookings").
		WithArgs(1, 10).
		WillReturnRows(rows)

	availableSeats, err := repo.GetAvailableSeats(context.Background(), 1, 10)

	assert.NoError(t, err)
	assert.Len(t, availableSeats, 2)
	assert.True(t, availableSeats[0].IsBooked)
	assert.True(t, availableSeats[1].IsBooked)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSeatRepository_GetAvailableSeats_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewSeatRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM seats s LEFT JOIN bookings").
		WithArgs(1, 10).
		WillReturnError(assert.AnError)

	availableSeats, err := repo.GetAvailableSeats(context.Background(), 1, 10)

	assert.Error(t, err)
	assert.Nil(t, availableSeats)
	assert.NoError(t, mock.ExpectationsWereMet())
}
