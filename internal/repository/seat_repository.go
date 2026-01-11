package repository

import (
	"context"
	"fmt"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SeatRepository interface {
	GetByCinemaID(ctx context.Context, cinemaID int) ([]*domain.Seat, error)
	GetByID(ctx context.Context, id int) (*domain.Seat, error)
	GetAvailableSeats(ctx context.Context, cinemaID, showtimeID int) ([]*domain.SeatAvailability, error)
}

type seatRepository struct {
	db *pgxpool.Pool
}

func NewSeatRepository(db *pgxpool.Pool) SeatRepository {
	return &seatRepository{db: db}
}

func (r *seatRepository) GetByCinemaID(ctx context.Context, cinemaID int) ([]*domain.Seat, error) {
	query := `
		SELECT id, cinema_id, seat_row, seat_number, seat_type, created_at
		FROM seats
		WHERE cinema_id = $1
		ORDER BY seat_row, seat_number
	`

	rows, err := r.db.Query(ctx, query, cinemaID)
	if err != nil {
		return nil, fmt.Errorf("failed to get seats: %w", err)
	}
	defer rows.Close()

	var seats []*domain.Seat
	for rows.Next() {
		var seat domain.Seat
		err := rows.Scan(
			&seat.ID,
			&seat.CinemaID,
			&seat.SeatRow,
			&seat.SeatNumber,
			&seat.SeatType,
			&seat.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan seat: %w", err)
		}
		seats = append(seats, &seat)
	}

	return seats, nil
}

func (r *seatRepository) GetByID(ctx context.Context, id int) (*domain.Seat, error) {
	query := `
		SELECT id, cinema_id, seat_row, seat_number, seat_type, created_at
		FROM seats
		WHERE id = $1
	`

	var seat domain.Seat
	err := r.db.QueryRow(ctx, query, id).Scan(
		&seat.ID,
		&seat.CinemaID,
		&seat.SeatRow,
		&seat.SeatNumber,
		&seat.SeatType,
		&seat.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("seat not found")
	}

	return &seat, nil
}

func (r *seatRepository) GetAvailableSeats(ctx context.Context, cinemaID, showtimeID int) ([]*domain.SeatAvailability, error) {
	query := `
		SELECT
			s.id, s.cinema_id, s.seat_row, s.seat_number, s.seat_type, s.created_at,
			CASE WHEN b.id IS NOT NULL THEN true ELSE false END as is_booked
		FROM seats s
		LEFT JOIN bookings b ON s.id = b.seat_id
			AND b.showtime_id = $2
			AND b.status IN ('pending', 'confirmed')
		WHERE s.cinema_id = $1
		ORDER BY s.seat_row, s.seat_number
	`

	rows, err := r.db.Query(ctx, query, cinemaID, showtimeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get seat availability: %w", err)
	}
	defer rows.Close()

	var seatAvailability []*domain.SeatAvailability
	for rows.Next() {
		var seat domain.Seat
		var isBooked bool

		err := rows.Scan(
			&seat.ID,
			&seat.CinemaID,
			&seat.SeatRow,
			&seat.SeatNumber,
			&seat.SeatType,
			&seat.CreatedAt,
			&isBooked,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan seat availability: %w", err)
		}

		seatAvailability = append(seatAvailability, &domain.SeatAvailability{
			Seat:       &seat,
			IsBooked:   isBooked,
			ShowtimeID: showtimeID,
		})
	}

	return seatAvailability, nil
}
