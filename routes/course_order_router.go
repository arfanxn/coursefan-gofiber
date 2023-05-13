package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerCourseOrderRouter registers courseOrder module routes into the router
func registerCourseOrderRouter(router fiber.Router) {
	courseOrderController := controllerp.InitCourseOrderController(databasep.MustGetGormDB())
	router.Get("/users/self/course-orders", courseOrderController.AllByAuthUser)

	courseOrders := router.Group("/course-orders")
	courseOrdersId := courseOrders.Group("/:course_order_id")
	courseOrdersId.Get("", courseOrderController.Find)
	courseOrders.Post("", courseOrderController.CreateByAuthUser)
}
