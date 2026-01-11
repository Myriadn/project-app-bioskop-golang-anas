package handler

import (
	"net/http"
	"strconv"

	"project-app-bioskop-golang-homework-anas/internal/service"
	"project-app-bioskop-golang-homework-anas/internal/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type SeatHandler struct {
	seatService service.SeatService
	logger      *zap.Logger
}

func NewSeatHandler(seatService service.SeatService, logger *zap.Logger) *SeatHandler {
	return &SeatHandler{
		seatService: seatService,
		logger:      logger,
	}
}

// Get seat availability for a specific cinema, date, and time
func (h *SeatHandler) GetSeatAvailability(w http.ResponseWriter, r *http.Request) {
	// Get cinema ID from URL parameter
	cinemaIDStr := chi.URLParam(r, "cinemaId")
	cinemaID, err := strconv.Atoi(cinemaIDStr)
	if err != nil {
		h.logger.Error("Invalid cinema ID", zap.String("cinema_id", cinemaIDStr), zap.Error(err))
		utils.SendBadRequest(w, "Invalid cinema ID", err)
		return
	}

	// Get date and time from query parameters
	date := r.URL.Query().Get("date")
	time := r.URL.Query().Get("time")

	if date == "" || time == "" {
		h.logger.Error("Missing date or time parameter")
		utils.SendBadRequest(w, "Date and time parameters are required", nil)
		return
	}

	// Get seat availability
	seats, showtime, err := h.seatService.GetSeatAvailability(r.Context(), cinemaID, date, time)
	if err != nil {
		h.logger.Error("Failed to get seat availability",
			zap.Int("cinema_id", cinemaID),
			zap.String("date", date),
			zap.String("time", time),
			zap.Error(err),
		)
		utils.SendNotFound(w, err.Error())
		return
	}

	h.logger.Info("Seat availability retrieved successfully",
		zap.Int("cinema_id", cinemaID),
		zap.String("date", date),
		zap.String("time", time),
		zap.Int("total_seats", len(seats)),
	)

	// Create response with showtime info and seats
	response := map[string]interface{}{
		"showtime": showtime,
		"seats":    seats,
	}

	utils.SendSuccess(w, "Seat availability retrieved successfully", response)
}
