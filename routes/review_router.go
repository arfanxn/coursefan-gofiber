package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerReviewRouter registers lecture module routes into the router
func registerReviewRouter(router fiber.Router) {
	reviewController := controllerp.InitReviewController(databasep.MustGetGormDB())
	coursesIdReviews := router.Group("/courses/:course_id/reviews")
	coursesIdReviews.Get("", reviewController.AllByCourse)
	coursesIdReviews.Post("", reviewController.CreateByCourse)

	reviews := router.Group("/reviews")
	reviewsId := reviews.Group("/:review_id")
	reviews.Post("", reviewController.Create)
	reviewsId.Get("", reviewController.Find)
	reviewsId.Put("", reviewController.Update)
	reviewsId.Delete("", reviewController.Delete)
}
