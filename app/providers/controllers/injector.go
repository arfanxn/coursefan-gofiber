//go:build wireinject
// +build wireinject

package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/controllers"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func initAuthController(db *gorm.DB) *controllers.AuthController {
	wire.Build(repositories.NewUserRepository, services.NewAuthService, controllers.NewAuthController)
	return nil
}
