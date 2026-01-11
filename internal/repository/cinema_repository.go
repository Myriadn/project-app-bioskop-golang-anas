package repository

import (
	"context"
	"errors"
	"fmt"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5"
)

type CinemaRepository interface {
	GetAll(ctx context.Context, limit, offset int) ([]*domain.Cinema, int, error)
	GetByID(ctx context.Context, id int) (*domain.Cinema, error)
}

type cinemaRepository struct {
	db PgxPool
}

func NewCinemaRepository(db PgxPool) CinemaRepository {
	return &cinemaRepository{db: db}
}

func (r *cinemaRepository) GetAll(ctx context.Context, limit, offset int) ([]*domain.Cinema, int, error) {
	// Count total
	var total int
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM cinemas").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count cinemas: %w", err)
	}

	// Get cinemas with pagination
	query := `
		SELECT id, name, location, description, created_at
		FROM cinemas
		ORDER BY id ASC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get cinemas: %w", err)
	}
	defer rows.Close()

	var cinemas []*domain.Cinema
	for rows.Next() {
		var cinema domain.Cinema
		err := rows.Scan(
			&cinema.ID,
			&cinema.Name,
			&cinema.Location,
			&cinema.Description,
			&cinema.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan cinema: %w", err)
		}
		cinemas = append(cinemas, &cinema)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating cinemas: %w", err)
	}

	return cinemas, total, nil
}

func (r *cinemaRepository) GetByID(ctx context.Context, id int) (*domain.Cinema, error) {
	query := `
		SELECT id, name, location, description, created_at
		FROM cinemas
		WHERE id = $1
	`

	var cinema domain.Cinema
	err := r.db.QueryRow(ctx, query, id).Scan(
		&cinema.ID,
		&cinema.Name,
		&cinema.Location,
		&cinema.Description,
		&cinema.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("cinema not found")
		}
		return nil, fmt.Errorf("failed to get cinema: %w", err)
	}

	return &cinema, nil
}
