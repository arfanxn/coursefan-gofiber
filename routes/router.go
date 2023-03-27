package routes

import "github.com/gofiber/fiber/v2"

// InjectRoutes will inject routes into application instance
func InjectRoutes(app *fiber.App) error {

	api := app.Group("/api")
	registerAuthRoutes(api)

	return nil
}
