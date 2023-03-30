package routes

import (
	"github.com/gofiber/fiber/v2"
)

// InjectRoutes will inject routes into application instance
func InjectRoutes(app *fiber.App) {

	router := app.Group("")
	registerMainMiddlewares(router)

	router.Static("/public", "./public") // File server route

	api := router.Group("/api")

	registerAuthRoutes(api)
}
