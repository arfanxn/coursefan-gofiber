package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerCourseRouter registers course module routes into the router
func registerCourseRouter(router fiber.Router) {
	courseController := controllerp.InitCourseController(databasep.MustGetGormDB())
	courses := router.Group("/courses")
	courses.Get("", courseController.All)
	courses.Get("/:course_id", courseController.Find)
	courses.Post("", courseController.Create)
	courses.Put("/:course_id", courseController.Update)
	courses.Delete("/:course_id", courseController.Delete)
}
