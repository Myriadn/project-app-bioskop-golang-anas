package handler

import (
	"encoding/json"
	"net/http"

	"project-app-bioskop-golang-homework-anas/internal/domain"
	"project-app-bioskop-golang-homework-anas/internal/service"
	"project-app-bioskop-golang-homework-anas/internal/utils"
	"project-app-bioskop-golang-homework-anas/pkg/validator"

	"go.uber.org/zap"
)

type PaymentHandler struct {
	paymentService service.PaymentService
	logger         *zap.Logger
}

func NewPaymentHandler(paymentService service.PaymentService, logger *zap.Logger) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		logger:         logger,
	}
}

// ProcessPayment godoc
// @Summary Process payment for booking
// @Description Process payment for an existing booking
// @Tags payments
// @Accept json
// @Produce json
// @Param request body domain.PaymentRequest true "Payment Request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/pay [post]
func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req domain.PaymentRequest

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

	// Process payment
	payment, err := h.paymentService.ProcessPayment(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to process payment",
			zap.Int("booking_id", req.BookingID),
			zap.Error(err),
		)
		utils.SendBadRequest(w, err.Error(), nil)
		return
	}

	h.logger.Info("Payment processed successfully",
		zap.Int("payment_id", payment.ID),
		zap.Int("booking_id", req.BookingID),
		zap.Float64("amount", payment.Amount),
	)

	utils.SendSuccess(w, "Payment processed successfully", payment)
}
