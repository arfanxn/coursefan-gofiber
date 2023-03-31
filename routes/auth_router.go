package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/database"
	"github.com/gofiber/fiber/v2"
)

// registerAuthRouter registers auth module routes into the router
func registerAuthRouter(router fiber.Router) {
	authController := controllerp.InitAuthController(database.MustGetGormDB())
	users := router.Group("/users")
	users.Post("/login", authController.Login).Name("login")
	users.Delete("/logout", authController.Logout).Name("logout")
	users.Post("/register", authController.Register).Name("register")
	users.Post("/forgot-password", authController.ForgotPassword).Name("forgot-password")
	users.Post("/reset-password", authController.ResetPassword).Name("reset-password")
}
