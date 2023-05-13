package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerWebhookRouter registers all webhook routes into the router
func registerWebhookRouter(router fiber.Router) {
	courseOrderController := controllerp.InitCourseOrderController(databasep.MustGetGormDB())
	webhooks := router.Group("/webhooks")
	webhooks.Post(
		"/course-orders/update-by-midtrans-notification",
		courseOrderController.UpdateByMidtransNotification,
	)

}
