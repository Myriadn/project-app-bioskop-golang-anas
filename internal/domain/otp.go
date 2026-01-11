package domain

import "time"

type OTPCode struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Code      string    `json:"code" db:"code"`
	IsUsed    bool      `json:"is_used" db:"is_used"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Request DTO
type VerifyOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,len=6"`
}

// Response DTO
type ResendOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
}
