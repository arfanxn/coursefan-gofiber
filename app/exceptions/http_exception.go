package exceptions

import "github.com/gofiber/fiber/v2"

// HTTP exceptions
var (
	HTTP404NotFound = fiber.NewError(fiber.StatusNotFound, "404 - Not Found")
)
