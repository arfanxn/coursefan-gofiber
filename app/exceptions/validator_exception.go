package exceptions

import "github.com/gofiber/fiber/v2"

// Validator exceptions
var (
	ValidatorUnknownTranslator = fiber.NewError(fiber.StatusInternalServerError, "Validator unknown translator")
)
