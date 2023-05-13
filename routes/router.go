package routes

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterApp will inject routes into application instance
func RegisterApp(app *fiber.App) {

	router := app.Group("")
	registerMiddlewareRouter(router)
	registerSandboxRouter(router)

	router.Static("/public", "./public", fiber.Static{
		ByteRange: true,
	}) // File server route

	api := router.Group("/api")

	registerAuthRouter(api)
	registerUserRouter(api)
	registerWalletRouter(api)
	registerTransactionRouter(api)
	registerNotificationRouter(api)
	registerCourseRouter(api)
	registerCourseOrderRouter(api)
	registerLecturePartRouter(api)
	registerLectureRouter(api)
	registerReviewRouter(api)
	registerDiscussionRouter(api)
}
