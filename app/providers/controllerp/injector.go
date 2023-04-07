//go:build wireinject
// +build wireinject

// controllerp  stands for Controller Provider
package controllerp

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/controllers"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitAuthController(db *gorm.DB) *controllers.AuthController {
	wire.Build(
		repositories.NewUserRepository,
		repositories.NewMediaRepository,
		repositories.NewTokenRepository,
		services.NewAuthService,
		controllers.NewAuthController,
	)
	return nil
}

func InitCourseController(db *gorm.DB) *controllers.CourseController {
	wire.Build(
		repositories.NewCourseRepository,
		services.NewCourseService,
		policies.NewCoursePolicy,
		controllers.NewCourseController,
	)
	return nil
}
