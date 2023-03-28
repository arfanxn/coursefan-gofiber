package routes

import (
	controller_provider "github.com/arfanxn/coursefan-gofiber/app/providers/controllers"
	"github.com/arfanxn/coursefan-gofiber/database"
	"github.com/gofiber/fiber/v2"
)

func registerAuthRoutes(router fiber.Router) {
	authController := controller_provider.AuthController(database.MustGetGORM())
	router.Post("/login", authController.Login)
}
