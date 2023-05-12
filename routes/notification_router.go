package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerNotificationRouter registers notification module routes into the router
func registerNotificationRouter(router fiber.Router) {
	notificationController := controllerp.InitNotificationController(databasep.MustGetGormDB())

	usersSelfNotifications := router.Group("/users/self/notifications")
	usersSelfNotifications.Get("", notificationController.AllByAuthUser)
	usersSelfNotifications.Post("", notificationController.CreateByAuthUser)

	notifications := router.Group("/notifications")
	notifications.Get("/:notification_id", notificationController.Find)
	notifications.Put("/:notification_id", notificationController.Update)
	notifications.Patch("/:notification_id/read", notificationController.MarkRead)
	notifications.Patch("/:notification_id/unread", notificationController.MarkUnread)
	notifications.Delete("/:notification_id", notificationController.Delete)
}
