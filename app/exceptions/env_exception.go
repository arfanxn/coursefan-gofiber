package exceptions

import "github.com/gofiber/fiber/v2"

// Environment exceptions
var (
	ENVVariableIsNotSetOrInvalid = fiber.NewError(fiber.StatusInternalServerError, "Environment variable is not set or has a invalid value")
)
