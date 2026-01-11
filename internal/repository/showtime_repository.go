package repository

import (
	"context"
	"fmt"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5"
)

type ShowtimeRepository interface {
	GetByCinemaDateTime(ctx context.Context, cinemaID int, date, time string) (*domain.Showtime, error)
	GetByID(ctx context.Context, id int) (*domain.Showtime, error)
}

type showtimeRepository struct {
	db PgxPool
}

func NewShowtimeRepository(db PgxPool) ShowtimeRepository {
	return &showtimeRepository{db: db}
}

func (r *showtimeRepository) GetByCinemaDateTime(ctx context.Context, cinemaID int, date, time string) (*domain.Showtime, error) {
	// Query dengan casting parameter ke DATE dan TIME
	query := `
		SELECT s.id, s.cinema_id, s.movie_id, s.show_date, s.show_time, s.price, s.created_at,
		       c.id, c.name, c.location, c.description, c.created_at,
		       m.id, m.title, m.description, m.duration, m.genre, m.poster_url, m.rating, m.created_at
		FROM showtimes s
		JOIN cinemas c ON s.cinema_id = c.id
		JOIN movies m ON s.movie_id = m.id
		WHERE s.cinema_id = $1
		  AND s.show_date = $2::date
		  AND s.show_time = $3::time
	`

	var showtime domain.Showtime
	var cinema domain.Cinema
	var movie domain.Movie

	err := r.db.QueryRow(ctx, query, cinemaID, date, time).Scan(
		&showtime.ID,
		&showtime.CinemaID,
		&showtime.MovieID,
		&showtime.ShowDate,
		&showtime.ShowTime,
		&showtime.Price,
		&showtime.CreatedAt,
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
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("showtime not found for cinema_id=%d, date=%s, time=%s", cinemaID, date, time)
		}
		return nil, fmt.Errorf("failed to get showtime: %w", err)
	}

	showtime.Cinema = &cinema
	showtime.Movie = &movie

	return &showtime, nil
}

func (r *showtimeRepository) GetByID(ctx context.Context, id int) (*domain.Showtime, error) {
	query := `
		SELECT s.id, s.cinema_id, s.movie_id, s.show_date, s.show_time, s.price, s.created_at,
		       c.id, c.name, c.location, c.description, c.created_at,
		       m.id, m.title, m.description, m.duration, m.genre, m.poster_url, m.rating, m.created_at
		FROM showtimes s
		JOIN cinemas c ON s.cinema_id = c.id
		JOIN movies m ON s.movie_id = m.id
		WHERE s.id = $1
	`

	var showtime domain.Showtime
	var cinema domain.Cinema
	var movie domain.Movie

	err := r.db.QueryRow(ctx, query, id).Scan(
		&showtime.ID,
		&showtime.CinemaID,
		&showtime.MovieID,
		&showtime.ShowDate,
		&showtime.ShowTime,
		&showtime.Price,
		&showtime.CreatedAt,
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
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("showtime not found")
		}
		return nil, fmt.Errorf("failed to get showtime: %w", err)
	}

	showtime.Cinema = &cinema
	showtime.Movie = &movie

	return &showtime, nil
}
