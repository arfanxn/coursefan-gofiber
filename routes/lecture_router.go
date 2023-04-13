package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerLectureRouter registers lecture module routes into the router
func registerLectureRouter(router fiber.Router) {
	lectureController := controllerp.InitLectureController(databasep.MustGetGormDB())
	coursesIdLectureParts := router.Group("/lecture_parts/:lecture_part_id/lectures")
	coursesIdLectureParts.Get("", lectureController.AllByLecturePart)
	coursesIdLectureParts.Get("/:lecture_id", lectureController.Find)
	coursesIdLectureParts.Post("", lectureController.Create)
	coursesIdLectureParts.Put("/:lecture_id", lectureController.Update)
	coursesIdLectureParts.Delete("/:lecture_id", lectureController.Delete)
}
