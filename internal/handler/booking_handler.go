package handler

import (
	"encoding/json"
	"net/http"

	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/middleware"
	"project-app-bioskop-golang-homework-anas/internal/service"
	"project-app-bioskop-golang-homework-anas/internal/utils"
	"project-app-bioskop-golang-homework-anas/pkg/validator"

	"go.uber.org/zap"
)

type BookingHandler struct {
	bookingService service.BookingService
	logger         *zap.Logger
}

func NewBookingHandler(bookingService service.BookingService, logger *zap.Logger) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
		logger:         logger,
	}
}

// Create a new ticket booking for a specific showtime and seat
func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Get user from context (set by auth middleware)
	user, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		h.logger.Error("User not found in context")
		utils.SendUnauthorized(w, "Unauthorized")
		return
	}

	var req domain.BookingRequest

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request", zap.Error(err))
		utils.SendBadRequest(w, "Invalid request body", err)
		return
	}

	// Validate request
	if err := validator.ValidateStruct(&req); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		utils.SendBadRequest(w, "Validation failed", err)
		return
	}

	// Create booking
	booking, err := h.bookingService.CreateBooking(r.Context(), user.ID, &req)
	if err != nil {
		h.logger.Error("Failed to create booking",
			zap.Int("user_id", user.ID),
			zap.Error(err),
		)
		utils.SendBadRequest(w, err.Error(), nil)
		return
	}

	h.logger.Info("Booking created successfully",
		zap.Int("booking_id", booking.ID),
		zap.Int("user_id", user.ID),
		zap.String("booking_code", booking.BookingCode),
	)

	utils.SendCreated(w, "Booking created successfully", booking)
}

// Get all bookings for the authenticated user
func (h *BookingHandler) GetUserBookings(w http.ResponseWriter, r *http.Request) {
	// Get user from context (set by auth middleware)
	user, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		h.logger.Error("User not found in context")
		utils.SendUnauthorized(w, "Unauthorized")
		return
	}

	// Get user bookings
	bookings, err := h.bookingService.GetUserBookings(r.Context(), user.ID)
	if err != nil {
		h.logger.Error("Failed to get user bookings",
			zap.Int("user_id", user.ID),
			zap.Error(err),
		)
		utils.SendInternalServerError(w, "Failed to get bookings", err)
		return
	}

	h.logger.Info("User bookings retrieved successfully",
		zap.Int("user_id", user.ID),
		zap.Int("total_bookings", len(bookings)),
	)

	utils.SendSuccess(w, "Bookings retrieved successfully", bookings)
}
