package routes

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterApp will inject routes into application instance
func RegisterApp(app *fiber.App) {

	router := app.Group("")
	registerMiddlewareRouter(router)

	router.Static("/public", "./public") // File server route

	api := router.Group("/api")

	registerAuthRouter(api)
}
