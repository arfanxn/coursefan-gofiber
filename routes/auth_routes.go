package routes

import (
	controller_provider "github.com/arfanxn/coursefan-gofiber/app/providers/controllers"
	"github.com/gofiber/fiber/v2"
)

func registerAuthRoutes(router fiber.Router) {
	authController := controller_provider.GetAuthController()
	users := router.Group("/users")
	users.Post("/login", authController.Login).Name("login")
	users.Post("/register", authController.Register).Name("register")
	users.Post("/forgot-password", authController.ForgotPassword).Name("forgot-password")
	users.Post("/reset-password", authController.ResetPassword).Name("reset-password")
}
