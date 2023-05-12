package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerUserRouter registers user module routes into the router
func registerUserRouter(router fiber.Router) {
	userController := controllerp.InitUserController(databasep.MustGetGormDB())
	userSettingController := controllerp.InitUserSettingController(databasep.MustGetGormDB())

	// User routes
	router.Get("/courses/:course_id/users", userController.AllByCourse)
	users := router.Group("/users")
	users.Get("/:user_id", userController.Find)
	usersSelf := users.Group("/self")
	usersSelf.Get("", userController.Self)
	usersSelf.Put("", userController.SelfUpdate)

	// User setting routes
	usersSelfSettings := usersSelf.Group("/settings")
	usersSelfSettings.Get("", userSettingController.AllByAuthUser)
	usersSelfSettings.Put("/:user_setting_key", userSettingController.UpdateByAuthUser)
}
