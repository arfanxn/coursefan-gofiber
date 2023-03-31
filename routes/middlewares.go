package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

// registerMiddlewareRouter registers main middlewares (required middlewares) to the router
func registerMiddlewareRouter(router fiber.Router) {
	router.Use(
		middlewares.Recovery(),
		middlewares.Language(),
		middlewares.Auth(),
		middlewares.After(), // after middleware
	)
}
