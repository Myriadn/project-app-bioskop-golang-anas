package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
)

// LogToFileAsync menulis log ke file secara async menggunakan goroutine
func LogToFileAsync(logger *zap.Logger, logDir string, message string, data map[string]interface{}) {
	go func() {
		// Buat directory jika belum ada
		if err := os.MkdirAll(logDir, 0755); err != nil {
			logger.Error("Failed to create log directory", zap.Error(err))
			return
		}

		// Generate filename dengan timestamp
		filename := filepath.Join(logDir, fmt.Sprintf("activity_%s.log", time.Now().Format("2006-01-02")))

		// Open file
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Error("Failed to open log file", zap.Error(err))
			return
		}
		defer file.Close()

		// Write log
		timestamp := time.Now().Format(time.RFC3339)
		logLine := fmt.Sprintf("[%s] %s - Data: %v\n", timestamp, message, data)

		if _, err := file.WriteString(logLine); err != nil {
			logger.Error("Failed to write to log file", zap.Error(err))
		}
	}()
}

// LogBookingAsync mencatat booking activity secara async
func LogBookingAsync(logger *zap.Logger, userID int, bookingCode string, action string) {
	go func() {
		data := map[string]interface{}{
			"user_id":      userID,
			"booking_code": bookingCode,
			"action":       action,
			"timestamp":    time.Now().Format(time.RFC3339),
		}

		LogToFileAsync(logger, "logs/bookings", fmt.Sprintf("Booking %s", action), data)
	}()
}

// LogPaymentAsync mencatat payment activity secara async
func LogPaymentAsync(logger *zap.Logger, bookingID int, amount float64, paymentMethod string) {
	go func() {
		data := map[string]interface{}{
			"booking_id":     bookingID,
			"amount":         amount,
			"payment_method": paymentMethod,
			"timestamp":      time.Now().Format(time.RFC3339),
		}

		LogToFileAsync(logger, "logs/payments", "Payment processed", data)
	}()
}
