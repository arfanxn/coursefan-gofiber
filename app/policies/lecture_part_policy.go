package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type LecturePartPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewLecturePartPolicy(
	permissionRepository *repositories.PermissionRepository,
) *LecturePartPolicy {
	return &LecturePartPolicy{
		permissionRepository: permissionRepository,
	}
}

// / AllByCourse policy ensures that the user has the right permissions for access a lecture parts.
func (policy *LecturePartPolicy) AllByCourse(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLecturePartView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Find policy ensures that the user has the right permissions for access a lecture part.
func (policy *LecturePartPolicy) Find(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLecturePartView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Create policy ensures that the user has the right permissions for create a lecture part.
func (policy *LecturePartPolicy) Create(c *fiber.Ctx, input requests.LecturePartCreate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLecturePartCreate)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Update policy ensures that the user has the right permissions for update a lecture part.
func (policy *LecturePartPolicy) Update(c *fiber.Ctx, input requests.LecturePartUpdate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLecturePartEdit)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Delete policy ensures that the user has the right permissions for delete a lecture part.
func (policy *LecturePartPolicy) Delete(c *fiber.Ctx, input requests.LecturePartDelete) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLecturePartDelete)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}
