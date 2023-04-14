package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type CoursePolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewCoursePolicy(permissionRepository *repositories.PermissionRepository) *CoursePolicy {
	return &CoursePolicy{permissionRepository: permissionRepository}
}

func (policy *CoursePolicy) Update(c *fiber.Ctx, input requests.CourseUpdate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameCourseEdit)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

func (policy *CoursePolicy) Delete(c *fiber.Ctx, input requests.CourseDelete) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameCourseDelete)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}
