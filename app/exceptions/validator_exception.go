package exceptions

import "github.com/gofiber/fiber/v2"

type ValidationError struct {
	Field, Message string
}

func (ValidationError *ValidationError) Error() string {
	return ValidationError.Message
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// Validator exceptions
var (
	ValidatorUnknownTranslator = fiber.NewError(fiber.StatusInternalServerError, "Validator unknown translator")
)
