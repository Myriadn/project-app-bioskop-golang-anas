package handler

import (
	"net/http"

	"project-app-bioskop-golang-homework-anas/internal/service"
	"project-app-bioskop-golang-homework-anas/internal/utils"

	"go.uber.org/zap"
)

type PaymentMethodHandler struct {
	paymentMethodService service.PaymentMethodService
	logger               *zap.Logger
}

func NewPaymentMethodHandler(paymentMethodService service.PaymentMethodService, logger *zap.Logger) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		paymentMethodService: paymentMethodService,
		logger:               logger,
	}
}

// Get list of all available payment methods
func (h *PaymentMethodHandler) GetAllPaymentMethods(w http.ResponseWriter, r *http.Request) {
	methods, err := h.paymentMethodService.GetAllPaymentMethods(r.Context())
	if err != nil {
		h.logger.Error("Failed to get payment methods", zap.Error(err))
		utils.SendInternalServerError(w, "Failed to get payment methods", err)
		return
	}

	h.logger.Info("Payment methods retrieved successfully", zap.Int("total", len(methods)))
	utils.SendSuccess(w, "Payment methods retrieved successfully", methods)
}
