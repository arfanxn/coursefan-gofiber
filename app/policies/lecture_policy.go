package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type LecturePolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewLecturePolicy(
	permissionRepository *repositories.PermissionRepository,
) *LecturePolicy {
	return &LecturePolicy{
		permissionRepository: permissionRepository,
	}
}

// AllByLecturePart policy ensures that the user has the right permissions for access a lectures.
func (policy *LecturePolicy) AllByLecturePart(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLectureView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Find policy ensures that the user has the right permissions for access a lecture part.
func (policy *LecturePolicy) Find(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLectureView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Create policy ensures that the user has the right permissions for create a lecture part.
func (policy *LecturePolicy) Create(c *fiber.Ctx, input requests.LectureCreate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLectureCreate)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Update policy ensures that the user has the right permissions for update a lecture part.
func (policy *LecturePolicy) Update(c *fiber.Ctx, input requests.LectureUpdate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLectureEdit)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Delete policy ensures that the user has the right permissions for delete a lecture part.
func (policy *LecturePolicy) Delete(c *fiber.Ctx, input requests.LectureDelete) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLectureDelete)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}
