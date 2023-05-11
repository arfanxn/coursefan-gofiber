package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerNotificationRouter registers notification module routes into the router
func registerNotificationRouter(router fiber.Router) {
	notificationController := controllerp.InitNotificationController(databasep.MustGetGormDB())
	notifications := router.Group("/notifications")
	notifications.Get("", notificationController.AllByAuthUser)
	notifications.Post("", notificationController.Create)
	notifications.Get("/:notification_id", notificationController.Find)
	notifications.Put("/:notification_id", notificationController.Update)
	notifications.Put("/:notification_id/read", notificationController.MarkRead)
	notifications.Put("/:notification_id/unread", notificationController.MarkUnread)
	notifications.Delete("/:notification_id", notificationController.Delete)
}
