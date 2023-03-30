package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/skip"
)

// registerMainMiddlewares registers main middlewares (required middlewares) to the router
func registerMainMiddlewares(router fiber.Router) {
	router.Use(
		middlewares.Recovery(),
		middlewares.Language(),
		skip.New(middlewares.Auth(), func(c *fiber.Ctx) bool {
			return false // skips if true
		}),
		middlewares.After(), // after middleware
	)
}
