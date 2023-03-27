package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func registerAuthRoutes(router fiber.Router) error {
	authController := controllers.NewAuthController()
	router.Post("/login", authController.Login)

	return nil
}
