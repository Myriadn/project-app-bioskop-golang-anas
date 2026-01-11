package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// InitValidator menginisialisasi validator
func InitValidator() {
	validate = validator.New()
}

// ValidateStruct memvalidasi struct
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// GetValidator mengembalikan validator instance
func GetValidator() *validator.Validate {
	return validate
}
