package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type UserPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewUserPolicy(permissionRepository *repositories.PermissionRepository) *UserPolicy {
	return &UserPolicy{permissionRepository: permissionRepository}
}

// AllByCourse policy ensures that the user has the right permissions for access a course users.
func (policy *UserPolicy) AllByCourse(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameCourseUserView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}
