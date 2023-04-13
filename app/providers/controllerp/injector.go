//go:build wireinject
// +build wireinject

package controllerp

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/controllers"
	"github.com/arfanxn/coursefan-gofiber/app/policies"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitAuthController(db *gorm.DB) *controllers.AuthController {
	wire.Build(repositories.NewUserRepository, repositories.NewTokenRepository, repositories.NewMediaRepository, services.NewAuthService, controllers.NewAuthController)
	return nil
}

func InitCourseController(db *gorm.DB) *controllers.CourseController {
	wire.Build(repositories.NewCourseRepository, repositories.NewCourseUserRoleRepository, services.NewCourseService, policies.NewCoursePolicy, controllers.NewCourseController)
	return nil
}

func InitLecturePartController(db *gorm.DB) *controllers.LecturePartController {
	wire.Build(repositories.NewCourseUserRoleRepository, repositories.NewLecturePartRepository, services.NewLecturePartService, policies.NewLecturePartPolicy, controllers.NewLecturePartController)
	return nil
}

func InitLectureController(db *gorm.DB) *controllers.LectureController {
	wire.Build(repositories.NewCourseUserRoleRepository, repositories.NewLectureRepository, repositories.NewMediaRepository, services.NewLectureService, policies.NewLecturePolicy, controllers.NewLectureController)
	return nil
}
