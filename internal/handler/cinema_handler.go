package handler

import (
	"net/http"
	"strconv"

	"project-app-bioskop-golang-homework-anas/internal/service"
	"project-app-bioskop-golang-homework-anas/internal/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CinemaHandler struct {
	cinemaService service.CinemaService
	logger        *zap.Logger
}

func NewCinemaHandler(cinemaService service.CinemaService, logger *zap.Logger) *CinemaHandler {
	return &CinemaHandler{
		cinemaService: cinemaService,
		logger:        logger,
	}
}

// GetAllCinemas godoc
// @Summary Get all cinemas
// @Description Get list of all cinemas with pagination
// @Tags cinemas
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Limit per page" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Failure 500 {object} utils.Response
// @Router /api/cinemas [get]
func (h *CinemaHandler) GetAllCinemas(w http.ResponseWriter, r *http.Request) {
	// Get pagination parameters from query
	page := 1
	limit := 10

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	// Get cinemas
	cinemas, meta, err := h.cinemaService.GetAllCinemas(r.Context(), page, limit)
	if err != nil {
		h.logger.Error("Failed to get cinemas", zap.Error(err))
		utils.SendInternalServerError(w, "Failed to get cinemas", err)
		return
	}

	h.logger.Info("Cinemas retrieved successfully",
		zap.Int("page", page),
		zap.Int("limit", limit),
		zap.Int("total", meta.TotalRows),
	)

	utils.SendPaginated(w, "Cinemas retrieved successfully", cinemas, meta)
}

// GetCinemaByID godoc
// @Summary Get cinema by ID
// @Description Get detailed information of a specific cinema
// @Tags cinemas
// @Accept json
// @Produce json
// @Param cinemaId path int true "Cinema ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /api/cinemas/{cinemaId} [get]
func (h *CinemaHandler) GetCinemaByID(w http.ResponseWriter, r *http.Request) {
	// Get cinema ID from URL parameter
	cinemaIDStr := chi.URLParam(r, "cinemaId")
	cinemaID, err := strconv.Atoi(cinemaIDStr)
	if err != nil {
		h.logger.Error("Invalid cinema ID", zap.String("cinema_id", cinemaIDStr), zap.Error(err))
		utils.SendBadRequest(w, "Invalid cinema ID", err)
		return
	}

	// Get cinema
	cinema, err := h.cinemaService.GetCinemaByID(r.Context(), cinemaID)
	if err != nil {
		h.logger.Error("Failed to get cinema", zap.Int("cinema_id", cinemaID), zap.Error(err))
		utils.SendNotFound(w, "Cinema not found")
		return
	}

	h.logger.Info("Cinema retrieved successfully", zap.Int("cinema_id", cinemaID))
	utils.SendSuccess(w, "Cinema retrieved successfully", cinema)
}
