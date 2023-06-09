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

func InitUserController(db *gorm.DB) *controllers.UserController {
	wire.Build(repositories.NewPermissionRepository,
		repositories.NewUserRepository,
		repositories.NewMediaRepository,
		repositories.NewUserProfileRepository,
		services.NewUserService,
		policies.NewUserPolicy,
		controllers.NewUserController,
	)
	return nil
}

func InitUserSettingController(db *gorm.DB) *controllers.UserSettingController {
	wire.Build(repositories.NewPermissionRepository, repositories.NewUserSettingRepository, services.NewUserSettingService, policies.NewUserSettingPolicy, controllers.NewUserSettingController)
	return nil
}

func InitAuthController(db *gorm.DB) *controllers.AuthController {
	wire.Build(repositories.NewUserRepository, repositories.NewTokenRepository, repositories.NewMediaRepository, services.NewAuthService, controllers.NewAuthController)
	return nil
}

func InitNotificationController(db *gorm.DB) *controllers.NotificationController {
	wire.Build(repositories.NewPermissionRepository, repositories.NewNotificationRepository, services.NewNotificationService, policies.NewNotificationPolicy, controllers.NewNotificationController)
	return nil
}

func InitCourseController(db *gorm.DB) *controllers.CourseController {
	wire.Build(repositories.NewCourseRepository, repositories.NewPermissionRepository, repositories.NewRoleRepository, repositories.NewCourseUserRoleRepository, services.NewCourseService, policies.NewCoursePolicy, controllers.NewCourseController)
	return nil
}

func InitCourseOrderController(db *gorm.DB) *controllers.CourseOrderController {
	wire.Build(repositories.NewCourseOrderRepository, repositories.NewCourseRepository, repositories.NewCourseUserRoleRepository, repositories.NewRoleRepository, repositories.NewPermissionRepository, services.NewCourseOrderService, policies.NewCourseOrderPolicy, controllers.NewCourseOrderController)
	return nil
}

func InitLecturePartController(db *gorm.DB) *controllers.LecturePartController {
	wire.Build(repositories.NewPermissionRepository, repositories.NewLecturePartRepository, services.NewLecturePartService, policies.NewLecturePartPolicy, controllers.NewLecturePartController)
	return nil
}

func InitLectureController(db *gorm.DB) *controllers.LectureController {
	wire.Build(repositories.NewPermissionRepository, repositories.NewLectureRepository, repositories.NewMediaRepository, services.NewLectureService, policies.NewLecturePolicy, controllers.NewLectureController)
	return nil
}

func InitReviewController(db *gorm.DB) *controllers.ReviewController {
	wire.Build(repositories.NewPermissionRepository, repositories.NewReviewRepository, services.NewReviewService, policies.NewReviewPolicy, controllers.NewReviewController)
	return nil
}

func InitDiscussionController(db *gorm.DB) *controllers.DiscussionController {
	wire.Build(repositories.NewPermissionRepository, repositories.NewDiscussionRepository, services.NewDiscussionService, policies.NewDiscussionPolicy, controllers.NewDiscussionController)
	return nil
}
