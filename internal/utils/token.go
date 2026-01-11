package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateToken menghasilkan random token
func GenerateToken(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateBookingCode menghasilkan booking code unik
func GenerateBookingCode() (string, error) {
	bytes := make([]byte, 6)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate booking code: %w", err)
	}
	return "BK" + hex.EncodeToString(bytes), nil
}
