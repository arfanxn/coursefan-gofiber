package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerDiscussionRouter registers discussion module routes into the router
func registerDiscussionRouter(router fiber.Router) {
	discussionController := controllerp.InitDiscussionController(databasep.MustGetGormDB())
	coursesIdDiscussions := router.Group("/lectures/:lecture_id/discussions")
	coursesIdDiscussions.Get("", discussionController.AllByLecture)
	coursesIdDiscussions.Post("", discussionController.CreateByLecture)

	discussions := router.Group("/discussions")
	discussionsId := discussions.Group("/:discussion_id")
	discussions.Post("", discussionController.Create)
	discussionsId.Get("", discussionController.Find)
	discussionsId.Put("", discussionController.Update)
	discussionsId.Delete("", discussionController.Delete)
}
