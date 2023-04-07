package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerLecturePartRouter registers course module routes into the router
func registerLecturePartRouter(router fiber.Router) {
	lecturePartController := controllerp.InitLecturePartController(databasep.MustGetGormDB())
	courses := router.Group("/courses")
	courses.Get("", lecturePartController.AllByCourse)
}
