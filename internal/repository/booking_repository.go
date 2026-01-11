package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *domain.Booking) error
	GetByID(ctx context.Context, id int) (*domain.Booking, error)
	GetByUserID(ctx context.Context, userID int) ([]*domain.Booking, error)
	Update(ctx context.Context, booking *domain.Booking) error
	CheckSeatBooked(ctx context.Context, showtimeID, seatID int) (bool, error)
}

type bookingRepository struct {
	db PgxPool
}

func NewBookingRepository(db PgxPool) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Create(ctx context.Context, booking *domain.Booking) error {
	query := `
		INSERT INTO bookings (user_id, showtime_id, seat_id, booking_code, status, total_price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	now := time.Now()
	err := r.db.QueryRow(
		ctx,
		query,
		booking.UserID,
		booking.ShowtimeID,
		booking.SeatID,
		booking.BookingCode,
		booking.Status,
		booking.TotalPrice,
		now,
		now,
	).Scan(&booking.ID, &booking.CreatedAt, &booking.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create booking: %w", err)
	}

	return nil
}

func (r *bookingRepository) GetByID(ctx context.Context, id int) (*domain.Booking, error) {
	query := `
		SELECT
			b.id, b.user_id, b.showtime_id, b.seat_id, b.booking_code, b.status, b.total_price, b.created_at, b.updated_at,
			s.id, s.cinema_id, s.movie_id, s.show_date, s.show_time, s.price, s.created_at,
			st.id, st.cinema_id, st.seat_row, st.seat_number, st.seat_type, st.created_at,
			c.id, c.name, c.location, c.description, c.created_at,
			m.id, m.title, m.description, m.duration, m.genre, m.poster_url, m.rating, m.created_at,
			p.id, p.booking_id, p.payment_method_id, p.amount, p.status, p.payment_details, p.paid_at, p.created_at,
			pm.id, pm.name, pm.code, pm.is_active, pm.created_at
		FROM bookings b
		JOIN showtimes s ON b.showtime_id = s.id
		JOIN seats st ON b.seat_id = st.id
		JOIN cinemas c ON s.cinema_id = c.id
		JOIN movies m ON s.movie_id = m.id
		LEFT JOIN payments p ON b.id = p.booking_id
		LEFT JOIN payment_methods pm ON p.payment_method_id = pm.id
		WHERE b.id = $1
	`

	var booking domain.Booking
	var showtime domain.Showtime
	var seat domain.Seat
	var cinema domain.Cinema
	var movie domain.Movie

	// Payment fields nullable (LEFT JOIN)
	var paymentID *int
	var paymentBookingID *int
	var paymentMethodID *int
	var paymentAmount *float64
	var paymentStatus *string
	var paymentDetails *domain.PaymentDetails
	var paymentPaidAt *time.Time
	var paymentCreatedAt *time.Time
	var pmID *int
	var pmName *string
	var pmCode *string
	var pmIsActive *bool
	var pmCreatedAt *time.Time

	err := r.db.QueryRow(ctx, query, id).Scan(
		&booking.ID,
		&booking.UserID,
		&booking.ShowtimeID,
		&booking.SeatID,
		&booking.BookingCode,
		&booking.Status,
		&booking.TotalPrice,
		&booking.CreatedAt,
		&booking.UpdatedAt,
		&showtime.ID,
		&showtime.CinemaID,
		&showtime.MovieID,
		&showtime.ShowDate,
		&showtime.ShowTime,
		&showtime.Price,
		&showtime.CreatedAt,
		&seat.ID,
		&seat.CinemaID,
		&seat.SeatRow,
		&seat.SeatNumber,
		&seat.SeatType,
		&seat.CreatedAt,
		&cinema.ID,
		&cinema.Name,
		&cinema.Location,
		&cinema.Description,
		&cinema.CreatedAt,
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Duration,
		&movie.Genre,
		&movie.PosterURL,
		&movie.Rating,
		&movie.CreatedAt,
		// Payment fields (nullable)
		&paymentID,
		&paymentBookingID,
		&paymentMethodID,
		&paymentAmount,
		&paymentStatus,
		&paymentDetails,
		&paymentPaidAt,
		&paymentCreatedAt,
		&pmID,
		&pmName,
		&pmCode,
		&pmIsActive,
		&pmCreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("booking not found")
		}
		return nil, fmt.Errorf("failed to get booking: %w", err)
	}

	showtime.Cinema = &cinema
	showtime.Movie = &movie
	booking.Showtime = &showtime
	booking.Seat = &seat

	// Build Payment object if exists
	if paymentID != nil {
		payment := &domain.Payment{
			ID:              *paymentID,
			BookingID:       *paymentBookingID,
			PaymentMethodID: *paymentMethodID,
			Amount:          *paymentAmount,
			Status:          *paymentStatus,
			PaidAt:          paymentPaidAt,
			CreatedAt:       *paymentCreatedAt,
		}

		if paymentDetails != nil {
			payment.PaymentDetails = *paymentDetails
		}

		// Build PaymentMethod if exists
		if pmID != nil {
			payment.PaymentMethod = &domain.PaymentMethod{
				ID:        *pmID,
				Name:      *pmName,
				Code:      *pmCode,
				IsActive:  *pmIsActive,
				CreatedAt: *pmCreatedAt,
			}
		}

		booking.Payment = payment
	}

	return &booking, nil
}

func (r *bookingRepository) GetByUserID(ctx context.Context, userID int) ([]*domain.Booking, error) {
	query := `
		SELECT
			b.id, b.user_id, b.showtime_id, b.seat_id, b.booking_code, b.status, b.total_price, b.created_at, b.updated_at,
			s.id, s.cinema_id, s.movie_id, s.show_date, s.show_time, s.price, s.created_at,
			st.id, st.cinema_id, st.seat_row, st.seat_number, st.seat_type, st.created_at,
			c.id, c.name, c.location, c.description, c.created_at,
			m.id, m.title, m.description, m.duration, m.genre, m.poster_url, m.rating, m.created_at,
			p.id, p.booking_id, p.payment_method_id, p.amount, p.status, p.payment_details, p.paid_at, p.created_at,
			pm.id, pm.name, pm.code, pm.is_active, pm.created_at
		FROM bookings b
		JOIN showtimes s ON b.showtime_id = s.id
		JOIN seats st ON b.seat_id = st.id
		JOIN cinemas c ON s.cinema_id = c.id
		JOIN movies m ON s.movie_id = m.id
		LEFT JOIN payments p ON b.id = p.booking_id
		LEFT JOIN payment_methods pm ON p.payment_method_id = pm.id
		WHERE b.user_id = $1
		ORDER BY b.created_at DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user bookings: %w", err)
	}
	defer rows.Close()

	var bookings []*domain.Booking
	for rows.Next() {
		var booking domain.Booking
		var showtime domain.Showtime
		var seat domain.Seat
		var cinema domain.Cinema
		var movie domain.Movie

		// Payment fields nullable
		var paymentID *int
		var paymentBookingID *int
		var paymentMethodID *int
		var paymentAmount *float64
		var paymentStatus *string
		var paymentDetails *domain.PaymentDetails
		var paymentPaidAt *time.Time
		var paymentCreatedAt *time.Time
		var pmID *int
		var pmName *string
		var pmCode *string
		var pmIsActive *bool
		var pmCreatedAt *time.Time

		err := rows.Scan(
			&booking.ID,
			&booking.UserID,
			&booking.ShowtimeID,
			&booking.SeatID,
			&booking.BookingCode,
			&booking.Status,
			&booking.TotalPrice,
			&booking.CreatedAt,
			&booking.UpdatedAt,
			&showtime.ID,
			&showtime.CinemaID,
			&showtime.MovieID,
			&showtime.ShowDate,
			&showtime.ShowTime,
			&showtime.Price,
			&showtime.CreatedAt,
			&seat.ID,
			&seat.CinemaID,
			&seat.SeatRow,
			&seat.SeatNumber,
			&seat.SeatType,
			&seat.CreatedAt,
			&cinema.ID,
			&cinema.Name,
			&cinema.Location,
			&cinema.Description,
			&cinema.CreatedAt,
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Duration,
			&movie.Genre,
			&movie.PosterURL,
			&movie.Rating,
			&movie.CreatedAt,
			// Payment fields
			&paymentID,
			&paymentBookingID,
			&paymentMethodID,
			&paymentAmount,
			&paymentStatus,
			&paymentDetails,
			&paymentPaidAt,
			&paymentCreatedAt,
			&pmID,
			&pmName,
			&pmCode,
			&pmIsActive,
			&pmCreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan booking: %w", err)
		}

		showtime.Cinema = &cinema
		showtime.Movie = &movie
		booking.Showtime = &showtime
		booking.Seat = &seat

		// Build Payment object if exists
		if paymentID != nil {
			payment := &domain.Payment{
				ID:              *paymentID,
				BookingID:       *paymentBookingID,
				PaymentMethodID: *paymentMethodID,
				Amount:          *paymentAmount,
				Status:          *paymentStatus,
				PaidAt:          paymentPaidAt,
				CreatedAt:       *paymentCreatedAt,
			}

			if paymentDetails != nil {
				payment.PaymentDetails = *paymentDetails
			}

			// Build PaymentMethod if exists
			if pmID != nil {
				payment.PaymentMethod = &domain.PaymentMethod{
					ID:        *pmID,
					Name:      *pmName,
					Code:      *pmCode,
					IsActive:  *pmIsActive,
					CreatedAt: *pmCreatedAt,
				}
			}

			booking.Payment = payment
		}

		bookings = append(bookings, &booking)
	}

	return bookings, nil
}

func (r *bookingRepository) Update(ctx context.Context, booking *domain.Booking) error {
	query := `
		UPDATE bookings
		SET status = $1, updated_at = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(ctx, query, booking.Status, time.Now(), booking.ID)
	if err != nil {
		return fmt.Errorf("failed to update booking: %w", err)
	}

	return nil
}

func (r *bookingRepository) CheckSeatBooked(ctx context.Context, showtimeID, seatID int) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM bookings
			WHERE showtime_id = $1
			  AND seat_id = $2
			  AND status IN ('pending', 'confirmed')
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, showtimeID, seatID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check seat booking: %w", err)
	}

	return exists, nil
}
