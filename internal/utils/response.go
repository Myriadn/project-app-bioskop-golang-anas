package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

type PaginatedResponse struct {
	Success    bool            `json:"success"`
	Message    string          `json:"message"`
	Data       interface{}     `json:"data"`
	Pagination *PaginationMeta `json:"pagination"`
}

// SendJSON mengirim response JSON
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// SendSuccess mengirim response sukses
func SendSuccess(w http.ResponseWriter, message string, data interface{}) {
	SendJSON(w, http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendCreated mengirim response created (201)
func SendCreated(w http.ResponseWriter, message string, data interface{}) {
	SendJSON(w, http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendError mengirim response error
func SendError(w http.ResponseWriter, statusCode int, message string, err error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	SendJSON(w, statusCode, Response{
		Success: false,
		Message: message,
		Error:   errMsg,
	})
}

// SendBadRequest mengirim response bad request (400)
func SendBadRequest(w http.ResponseWriter, message string, err error) {
	SendError(w, http.StatusBadRequest, message, err)
}

// SendUnauthorized mengirim response unauthorized (401)
func SendUnauthorized(w http.ResponseWriter, message string) {
	SendError(w, http.StatusUnauthorized, message, nil)
}

// SendNotFound mengirim response not found (404)
func SendNotFound(w http.ResponseWriter, message string) {
	SendError(w, http.StatusNotFound, message, nil)
}

// SendInternalServerError mengirim response internal server error (500)
func SendInternalServerError(w http.ResponseWriter, message string, err error) {
	SendError(w, http.StatusInternalServerError, message, err)
}

// SendPaginated mengirim response dengan pagination
func SendPaginated(w http.ResponseWriter, message string, data interface{}, meta *PaginationMeta) {
	SendJSON(w, http.StatusOK, PaginatedResponse{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: meta,
	})
}
