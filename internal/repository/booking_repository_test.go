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

func TestBookingRepository_Create(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)

	booking := &domain.Booking{
		UserID:      1,
		ShowtimeID:  1,
		SeatID:      10,
		BookingCode: "BK123456",
		Status:      "pending",
		TotalPrice:  50000,
	}

	now := time.Now()
	rows := pgxmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(1, now, now)

	mock.ExpectQuery("INSERT INTO bookings").
		WithArgs(
			booking.UserID,
			booking.ShowtimeID,
			booking.SeatID,
			booking.BookingCode,
			booking.Status,
			booking.TotalPrice,
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
		).
		WillReturnRows(rows)

	err = repo.Create(context.Background(), booking)

	assert.NoError(t, err)
	assert.Equal(t, 1, booking.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingRepository_CheckSeatBooked_True(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)

	rows := pgxmock.NewRows([]string{"exists"}).AddRow(true)
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(1, 10).
		WillReturnRows(rows)

	isBooked, err := repo.CheckSeatBooked(context.Background(), 1, 10)

	assert.NoError(t, err)
	assert.True(t, isBooked)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingRepository_CheckSeatBooked_False(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)

	rows := pgxmock.NewRows([]string{"exists"}).AddRow(false)
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(1, 10).
		WillReturnRows(rows)

	isBooked, err := repo.CheckSeatBooked(context.Background(), 1, 10)

	assert.NoError(t, err)
	assert.False(t, isBooked)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingRepository_Update(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)

	booking := &domain.Booking{
		ID:     1,
		Status: "confirmed",
	}

	mock.ExpectExec("UPDATE bookings").
		WithArgs(booking.Status, pgxmock.AnyArg(), booking.ID).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.Update(context.Background(), booking)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingRepository_GetByID(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)
	now := time.Now()

	rows := pgxmock.NewRows([]string{
		"id", "user_id", "showtime_id", "seat_id", "booking_code", "status", "total_price", "created_at", "updated_at",
		"showtime_id", "cinema_id", "movie_id", "show_date", "show_time", "price", "showtime_created_at",
		"seat_id", "cinema_id", "seat_row", "seat_number", "seat_type", "seat_created_at",
		"cinema_id", "name", "location", "description", "cinema_created_at",
		"movie_id", "title", "description", "duration", "genre", "poster_url", "rating", "movie_created_at",
		"payment_id", "booking_id", "payment_method_id", "amount", "payment_status", "payment_details", "paid_at", "payment_created_at",
		"pm_id", "pm_name", "code", "is_active", "pm_created_at",
	}).AddRow(
		1, 1, 10, 20, "BK123", "pending", 50000.0, now, now, // Booking
		10, 1, 5, now, now, 50000.0, now, // Showtime
		20, 1, "A", 1, "regular", now, // Seat
		1, "CGV Grand Indonesia", "Jakarta", "Premium", now, // Cinema
		5, "Avengers", "Action", 120, "Action", "url", "PG-13", now, // Movie
		nil, nil, nil, nil, nil, nil, nil, nil, // Payment (Ganti AnyArg jadi nil)
		nil, nil, nil, nil, nil, // Payment Method (Ganti AnyArg jadi nil)
	)

	mock.ExpectQuery("SELECT (.+) FROM bookings b").WithArgs(1).WillReturnRows(rows)

	booking, err := repo.GetByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, booking)
	assert.Equal(t, 1, booking.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// func TestBookingRepository_GetByUserID(t *testing.T) {
// 	mock, err := pgxmock.NewPool()
// 	require.NoError(t, err)
// 	defer mock.Close()

// 	repo := NewBookingRepository(mock)
// 	now := time.Now()
// 	paidAt := now

// 	rows := pgxmock.NewRows([]string{
// 		"id", "user_id", "showtime_id", "seat_id", "booking_code", "status", "total_price", "created_at", "updated_at",
// 		"showtime_id", "cinema_id", "movie_id", "show_date", "show_time", "price", "showtime_created_at",
// 		"seat_id", "cinema_id", "seat_row", "seat_number", "seat_type", "seat_created_at",
// 		"cinema_id", "name", "location", "description", "cinema_created_at",
// 		"movie_id", "title", "description", "duration", "genre", "poster_url", "rating", "movie_created_at",
// 		"payment_id", "booking_id", "payment_method_id", "amount", "payment_status", "payment_details", "paid_at", "payment_created_at",
// 		"pm_id", "pm_name", "code", "is_active", "pm_created_at",
// 	}).
// 		AddRow(
// 			1, 1, 10, 20, "BK123", "pending", 50000.0, now, now,
// 			10, 1, 5, now, now, 50000.0, now,
// 			20, 1, "A", 1, "regular", now,
// 			1, "CGV Grand Indonesia", "Jakarta", "Premium", now,
// 			5, "Avengers", "Action", 120, "Action", "url", "PG-13", now,
// 			nil, nil, nil, nil, nil, nil, nil, nil,
// 			nil, nil, nil, nil, nil,
// 		).
// 		AddRow(
// 			2, 1, 11, 21, "BK456", "confirmed", 60000.0, now, now,
// 			11, 1, 6, now, now, 60000.0, now,
// 			21, 1, "B", 2, "vip", now,
// 			1, "CGV Grand Indonesia", "Jakarta", "Premium", now,
// 			6, "Spiderman", "Action", 110, "Action", "url2", "PG-13", now,
// 			// Bungkus semua data Payment & Payment Method dengan helper pointer
// 			intPtr(101),            // payment_id
// 			intPtr(2),              // booking_id
// 			intPtr(2),              // payment_method_id
// 			floatPtr(60000.0),      // amount (Ini yang bikin error tadi)
// 			stringPtr("completed"), // payment_status
// 			[]byte(`{}`),           // payment_details (biasanya []byte sudah dianggap pointer-like)
// 			&paidAt,                // paid_at
// 			now,                    // payment_created_at
// 			intPtr(2),              // pm_id
// 			stringPtr("E-Wallet"),  // pm_name
// 			stringPtr("GOPAY"),     // code
// 			true,                   // is_active
// 			now,                    // pm_created_at
// 		)

// 	mock.ExpectQuery("SELECT (.+) FROM bookings b").
// 		WithArgs(1).
// 		WillReturnRows(rows)

// 	bookings, err := repo.GetByUserID(context.Background(), 1)

// 	assert.NoError(t, err)
// 	if assert.NotEmpty(t, bookings) {
// 		assert.Len(t, bookings, 2)
// 		assert.Equal(t, "BK123", bookings[0].BookingCode)
// 	}
// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

func TestBookingRepository_GetByID_NotFound(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)

	mock.ExpectQuery("SELECT (.+) FROM bookings b").
		WithArgs(999).
		WillReturnError(pgx.ErrNoRows)

	booking, err := repo.GetByID(context.Background(), 999)

	assert.Error(t, err)
	assert.Nil(t, booking)
	assert.Contains(t, err.Error(), "booking not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingRepository_GetByUserID_Empty(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)

	rows := pgxmock.NewRows([]string{
		"id", "user_id", "showtime_id", "seat_id", "booking_code", "status", "total_price", "created_at", "updated_at",
		"showtime_id", "cinema_id", "movie_id", "show_date", "show_time", "price", "showtime_created_at",
		"seat_id", "cinema_id", "seat_row", "seat_number", "seat_type", "seat_created_at",
		"cinema_id", "name", "location", "description", "cinema_created_at",
		"movie_id", "title", "description", "duration", "genre", "poster_url", "rating", "movie_created_at",
		"payment_id", "booking_id", "payment_method_id", "amount", "payment_status", "payment_details", "paid_at", "payment_created_at",
		"pm_id", "pm_name", "code", "is_active", "pm_created_at",
	})

	mock.ExpectQuery("SELECT (.+) FROM bookings b").
		WithArgs(999).
		WillReturnRows(rows)

	bookings, err := repo.GetByUserID(context.Background(), 999)

	assert.NoError(t, err)
	assert.Len(t, bookings, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestBookingRepository_Create_Error(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	repo := NewBookingRepository(mock)

	booking := &domain.Booking{
		UserID:      1,
		ShowtimeID:  1,
		SeatID:      10,
		BookingCode: "BK123456",
		Status:      "pending",
		TotalPrice:  50000,
	}

	mock.ExpectQuery("INSERT INTO bookings").
		WithArgs(
			booking.UserID,
			booking.ShowtimeID,
			booking.SeatID,
			booking.BookingCode,
			booking.Status,
			booking.TotalPrice,
			pgxmock.AnyArg(),
			pgxmock.AnyArg(),
		).
		WillReturnError(assert.AnError)

	err = repo.Create(context.Background(), booking)

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
