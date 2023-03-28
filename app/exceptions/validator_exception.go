package exceptions

import "github.com/gofiber/fiber/v2"

type ValidatorError struct {
	Field, Message string
}

func (validatorError *ValidatorError) Error() string {
	return validatorError.Message
}

func NewValidatorError(field, message string) *ValidatorError {
	return &ValidatorError{
		Field:   field,
		Message: message,
	}
}

// Validator exceptions
var (
	ValidatorUnknownTranslator = fiber.NewError(fiber.StatusInternalServerError, "Validator unknown translator")
)
