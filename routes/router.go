package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

// registerMainMiddlewares registers main middlewares (required middlewares) to the router
func registerMainMiddlewares(router fiber.Router) {
	router.Use(
		middlewares.Recovery(),
		middlewares.After(), // after middleware
	)
}

// InjectRoutes will inject routes into application instance
func InjectRoutes(app *fiber.App) error {

	api := app.Group("/api")
	registerMainMiddlewares(api)
	registerAuthRoutes(api)

	return nil
}
