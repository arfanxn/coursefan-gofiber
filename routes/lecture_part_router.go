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
	coursesIdLectureParts.Get("/:lecture_part_id", lecturePartController.Find)
	coursesIdLectureParts.Post("", lecturePartController.Create)
	coursesIdLectureParts.Put("/:lecture_part_id", lecturePartController.Update)
	coursesIdLectureParts.Delete("/:lecture_part_id", lecturePartController.Delete)
}
