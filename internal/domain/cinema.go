package domain

import "time"

type Cinema struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Location    string    `json:"location" db:"location"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Movie struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Duration    int       `json:"duration" db:"duration"` // dalam menit
	Genre       string    `json:"genre" db:"genre"`
	PosterURL   string    `json:"poster_url" db:"poster_url"`
	Rating      string    `json:"rating" db:"rating"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Showtime struct {
	ID        int       `json:"id" db:"id"`
	CinemaID  int       `json:"cinema_id" db:"cinema_id"`
	MovieID   int       `json:"movie_id" db:"movie_id"`
	ShowDate  string    `json:"show_date" db:"show_date"`
	ShowTime  string    `json:"show_time" db:"show_time"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// Relations
	Cinema *Cinema `json:"cinema,omitempty"`
	Movie  *Movie  `json:"movie,omitempty"`
}

type Seat struct {
	ID         int       `json:"id" db:"id"`
	CinemaID   int       `json:"cinema_id" db:"cinema_id"`
	SeatRow    string    `json:"seat_row" db:"seat_row"`
	SeatNumber int       `json:"seat_number" db:"seat_number"`
	SeatType   string    `json:"seat_type" db:"seat_type"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type SeatAvailability struct {
	Seat       *Seat `json:"seat"`
	IsBooked   bool  `json:"is_booked"`
	ShowtimeID int   `json:"showtime_id"`
}
