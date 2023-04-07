package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerLecturePartRouter registers lecture part module routes into the router
func registerLecturePartRouter(router fiber.Router) {
	lecturePartController := controllerp.InitLecturePartController(databasep.MustGetGormDB())
	coursesIdLectureParts := router.Group("/courses/:course_id/lecture_parts")
	coursesIdLectureParts.Get("", lecturePartController.AllByCourse)
}
