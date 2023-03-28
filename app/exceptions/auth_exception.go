package exceptions

import "github.com/gofiber/fiber/v2"

// Auth exceptions
var (
	AuthSessionExpired          = fiber.NewError(fiber.StatusUnauthorized, "Session expired please sign in")
	AuthCredentialsDoesNotMatch = fiber.NewError(fiber.StatusUnauthorized, "Credentials does not match our records")
)
