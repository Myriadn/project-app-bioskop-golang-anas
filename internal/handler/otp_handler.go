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

type OTPHandler struct {
	otpService service.OTPService
	logger     *zap.Logger
}

func NewOTPHandler(otpService service.OTPService, logger *zap.Logger) *OTPHandler {
	return &OTPHandler{
		otpService: otpService,
		logger:     logger,
	}
}

// VerifyOTP godoc
// @Summary Verify OTP code
// @Description Verify email using OTP code sent to user's email
// @Tags otp
// @Accept json
// @Produce json
// @Param request body domain.VerifyOTPRequest true "Verify OTP Request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/verify-otp [post]
func (h *OTPHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var req domain.VerifyOTPRequest

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

	// Verify OTP
	if err := h.otpService.VerifyOTP(r.Context(), req.Email, req.Code); err != nil {
		h.logger.Error("Failed to verify OTP",
			zap.String("email", req.Email),
			zap.Error(err),
		)
		utils.SendBadRequest(w, err.Error(), nil)
		return
	}

	h.logger.Info("OTP verified successfully", zap.String("email", req.Email))
	utils.SendSuccess(w, "Email verified successfully! You can now login.", nil)
}

// ResendOTP godoc
// @Summary Resend OTP code
// @Description Resend OTP code to user's email
// @Tags otp
// @Accept json
// @Produce json
// @Param request body domain.ResendOTPRequest true "Resend OTP Request"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/resend-otp [post]
func (h *OTPHandler) ResendOTP(w http.ResponseWriter, r *http.Request) {
	var req domain.ResendOTPRequest

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

	// Resend OTP
	if err := h.otpService.ResendOTP(r.Context(), req.Email); err != nil {
		h.logger.Error("Failed to resend OTP",
			zap.String("email", req.Email),
			zap.Error(err),
		)
		utils.SendBadRequest(w, err.Error(), nil)
		return
	}

	h.logger.Info("OTP resent successfully", zap.String("email", req.Email))
	utils.SendSuccess(w, "OTP code has been sent to your email", nil)
}
