package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/controllers"
	"github.com/arfanxn/coursefan-gofiber/database"
)

var authController *controllers.AuthController

func GetAuthController() *controllers.AuthController {
	if authController != nil {
		return authController
	}
	authController = initAuthController(database.MustGetGORM())
	return authController
}
