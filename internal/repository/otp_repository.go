package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"project-app-bioskop-golang-homework-anas/internal/domain"

	"github.com/jackc/pgx/v5"
)

type OTPRepository interface {
	Create(ctx context.Context, otp *domain.OTPCode) error
	GetByUserIDAndCode(ctx context.Context, userID int, code string) (*domain.OTPCode, error)
	MarkAsUsed(ctx context.Context, id int) error
	DeleteExpired(ctx context.Context) error
	DeleteByUserID(ctx context.Context, userID int) error
}

type otpRepository struct {
	db PgxPool
}

func NewOTPRepository(db PgxPool) OTPRepository {
	return &otpRepository{db: db}
}

func (r *otpRepository) Create(ctx context.Context, otp *domain.OTPCode) error {
	query := `
		INSERT INTO otp_codes (user_id, code, is_used, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	now := time.Now()
	err := r.db.QueryRow(
		ctx,
		query,
		otp.UserID,
		otp.Code,
		otp.IsUsed,
		otp.ExpiresAt,
		now,
	).Scan(&otp.ID, &otp.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create OTP: %w", err)
	}

	return nil
}

func (r *otpRepository) GetByUserIDAndCode(ctx context.Context, userID int, code string) (*domain.OTPCode, error) {
	query := `
		SELECT id, user_id, code, is_used, expires_at, created_at
		FROM otp_codes
		WHERE user_id = $1 AND code = $2 AND is_used = false AND expires_at > NOW()
		ORDER BY created_at DESC
		LIMIT 1
	`

	var otp domain.OTPCode
	err := r.db.QueryRow(ctx, query, userID, code).Scan(
		&otp.ID,
		&otp.UserID,
		&otp.Code,
		&otp.IsUsed,
		&otp.ExpiresAt,
		&otp.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("invalid or expired OTP")
		}
		return nil, fmt.Errorf("failed to get OTP: %w", err)
	}

	return &otp, nil
}

func (r *otpRepository) MarkAsUsed(ctx context.Context, id int) error {
	query := `UPDATE otp_codes SET is_used = true WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to mark OTP as used: %w", err)
	}

	return nil
}

func (r *otpRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM otp_codes WHERE expires_at <= NOW() OR is_used = true`

	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to delete expired OTPs: %w", err)
	}

	return nil
}

func (r *otpRepository) DeleteByUserID(ctx context.Context, userID int) error {
	query := `DELETE FROM otp_codes WHERE user_id = $1`

	_, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user OTPs: %w", err)
	}

	return nil
}
