package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthTokenRepository interface {
	Create(ctx context.Context, token *domain.AuthToken) error
	GetByToken(ctx context.Context, token string) (*domain.AuthToken, error)
	Delete(ctx context.Context, token string) error
	DeleteByUserID(ctx context.Context, userID int) error
	DeleteExpired(ctx context.Context) error
}

type authTokenRepository struct {
	db *pgxpool.Pool
}

func NewAuthTokenRepository(db *pgxpool.Pool) AuthTokenRepository {
	return &authTokenRepository{db: db}
}

func (r *authTokenRepository) Create(ctx context.Context, token *domain.AuthToken) error {
	query := `
		INSERT INTO auth_tokens (user_id, token, expires_at, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	now := time.Now()
	err := r.db.QueryRow(
		ctx,
		query,
		token.UserID,
		token.Token,
		token.ExpiresAt,
		now,
	).Scan(&token.ID, &token.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create auth token: %w", err)
	}

	return nil
}

func (r *authTokenRepository) GetByToken(ctx context.Context, token string) (*domain.AuthToken, error) {
	query := `
		SELECT id, user_id, token, expires_at, created_at
		FROM auth_tokens
		WHERE token = $1 AND expires_at > NOW()
	`

	var authToken domain.AuthToken
	err := r.db.QueryRow(ctx, query, token).Scan(
		&authToken.ID,
		&authToken.UserID,
		&authToken.Token,
		&authToken.ExpiresAt,
		&authToken.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("token not found or expired")
		}
		return nil, fmt.Errorf("failed to get auth token: %w", err)
	}

	return &authToken, nil
}

func (r *authTokenRepository) Delete(ctx context.Context, token string) error {
	query := `DELETE FROM auth_tokens WHERE token = $1`

	_, err := r.db.Exec(ctx, query, token)
	if err != nil {
		return fmt.Errorf("failed to delete auth token: %w", err)
	}

	return nil
}

func (r *authTokenRepository) DeleteByUserID(ctx context.Context, userID int) error {
	query := `DELETE FROM auth_tokens WHERE user_id = $1`

	_, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user tokens: %w", err)
	}

	return nil
}

func (r *authTokenRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM auth_tokens WHERE expires_at <= NOW()`

	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to delete expired tokens: %w", err)
	}

	return nil
}
