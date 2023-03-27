package exceptions

import "github.com/gofiber/fiber/v2"

// JWT exceptions
var (
	JWTExpired = fiber.NewError(fiber.StatusUnauthorized, "Token expired please sign in")
)
