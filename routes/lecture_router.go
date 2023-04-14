package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerLectureRouter registers lecture module routes into the router
func registerLectureRouter(router fiber.Router) {
	lectureController := controllerp.InitLectureController(databasep.MustGetGormDB())
	lecturePartsIdLectures := router.Group("/lecture_parts/:lecture_part_id/lectures")
	lecturePartsIdLectures.Get("", lectureController.AllByLecturePart)
	lecturePartsIdLectures.Get("/:lecture_id", lectureController.Find)
	lecturePartsIdLectures.Post("", lectureController.Create)
	lecturePartsIdLectures.Put("/:lecture_id", lectureController.Update)
	lecturePartsIdLectures.Delete("/:lecture_id", lectureController.Delete)
}
