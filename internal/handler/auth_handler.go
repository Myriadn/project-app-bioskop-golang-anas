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

type AuthHandler struct {
	authService service.AuthService
	logger      *zap.Logger
}

func NewAuthHandler(authService service.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// Register a new user account
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req domain.RegisterRequest

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

	// Register user
	authResp, err := h.authService.Register(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to register user", zap.Error(err))
		utils.SendBadRequest(w, err.Error(), nil)
		return
	}

	h.logger.Info("User registered successfully", zap.String("username", req.Username))
	utils.SendCreated(w, "User registered successfully", authResp)
}

// Login with username and password
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req domain.LoginRequest

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

	// Login user
	authResp, err := h.authService.Login(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to login", zap.Error(err))
		utils.SendUnauthorized(w, err.Error())
		return
	}

	h.logger.Info("User logged in successfully", zap.String("username", req.Username))
	utils.SendSuccess(w, "Login successful", authResp)
}

// Logout and invalidate token
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// Get token from header
	token := r.Header.Get("Authorization")
	if token == "" {
		utils.SendUnauthorized(w, "Authorization token required")
		return
	}

	// Remove "Bearer " prefix if exists
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Logout
	if err := h.authService.Logout(r.Context(), token); err != nil {
		h.logger.Error("Failed to logout", zap.Error(err))
		utils.SendInternalServerError(w, "Failed to logout", err)
		return
	}

	h.logger.Info("User logged out successfully")
	utils.SendSuccess(w, "Logout successful", nil)
}
