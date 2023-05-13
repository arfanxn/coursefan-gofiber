package routes

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterApp will inject routes into application instance
func RegisterApp(app *fiber.App) {
	router := app.Group("")

	// Register middleware to all routes
	registerMiddlewareRouter(router)

	// Sandbox routes
	registerSandboxRouter(router)

	// File server
	router.Static("/public", "./public", fiber.Static{
		ByteRange: true,
	}) // File server route

	// Sub route
	api := router.Group("/api")

	// Module routes
	registerAuthRouter(api)
	registerUserRouter(api)
	registerNotificationRouter(api)
	registerCourseRouter(api)
	registerCourseOrderRouter(api)
	registerLecturePartRouter(api)
	registerLectureRouter(api)
	registerReviewRouter(api)
	registerDiscussionRouter(api)

	// Webhook
	registerWebhookRouter(router)
}
