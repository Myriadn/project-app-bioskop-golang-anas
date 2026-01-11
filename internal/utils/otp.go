package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateOTP menghasilkan 6-digit OTP code
func GenerateOTP() (string, error) {
	const otpLength = 6
	const digits = "0123456789"

	otp := make([]byte, otpLength)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", fmt.Errorf("failed to generate OTP: %w", err)
		}
		otp[i] = digits[num.Int64()]
	}

	return string(otp), nil
}
