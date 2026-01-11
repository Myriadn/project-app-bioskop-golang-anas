package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Booking struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	ShowtimeID  int       `json:"showtime_id" db:"showtime_id"`
	SeatID      int       `json:"seat_id" db:"seat_id"`
	BookingCode string    `json:"booking_code" db:"booking_code"`
	Status      string    `json:"status" db:"status"`
	TotalPrice  float64   `json:"total_price" db:"total_price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	// Relations
	Showtime *Showtime `json:"showtime,omitempty"`
	Seat     *Seat     `json:"seat,omitempty"`
	Payment  *Payment  `json:"payment,omitempty"`
}

// PaymentDetails adalah custom type untuk handle JSONB
type PaymentDetails map[string]interface{}

// Value implements driver.Valuer interface
func (p PaymentDetails) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

// Scan implements sql.Scanner interface
func (p *PaymentDetails) Scan(value interface{}) error {
	if value == nil {
		*p = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	return json.Unmarshal(bytes, p)
}

type Payment struct {
	ID              int            `json:"id" db:"id"`
	BookingID       int            `json:"booking_id" db:"booking_id"`
	PaymentMethodID int            `json:"payment_method_id" db:"payment_method_id"`
	Amount          float64        `json:"amount" db:"amount"`
	Status          string         `json:"status" db:"status"`
	PaymentDetails  PaymentDetails `json:"payment_details,omitempty" db:"payment_details"`
	PaidAt          *time.Time     `json:"paid_at" db:"paid_at"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	// Relations
	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
}

type PaymentMethod struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Code      string    `json:"code" db:"code"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Request DTOs
type BookingRequest struct {
	CinemaID      int    `json:"cinema_id" validate:"required"`
	SeatID        int    `json:"seat_id" validate:"required"`
	Date          string `json:"date" validate:"required"`
	Time          string `json:"time" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type PaymentRequest struct {
	BookingID      int            `json:"booking_id" validate:"required"`
	PaymentMethod  string         `json:"payment_method" validate:"required"`
	PaymentDetails PaymentDetails `json:"payment_details"`
}
