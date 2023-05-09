package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type ReviewPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewReviewPolicy(permissionRepository *repositories.PermissionRepository) *ReviewPolicy {
	return &ReviewPolicy{permissionRepository: permissionRepository}
}

// AllByCourse policy ensures that the user has the right permissions for access a course reviews.
func (policy *ReviewPolicy) AllByCourse(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameCourseReviewView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Find policy ensures that the user has the right permissions for access a review.
func (policy *ReviewPolicy) Find(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameReviewView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Create policy ensures that the user has the right permissions for create a review.
func (policy *ReviewPolicy) Create(c *fiber.Ctx, input requests.ReviewCreate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameReviewCreate)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// CreateByCourse policy ensures that the user has the right permissions for create a review by the given course.
func (policy *ReviewPolicy) CreateByCourse(c *fiber.Ctx, input requests.ReviewCreate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameCourseReviewCreate)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Update policy ensures that the user has the right permissions for update a review.
func (policy *ReviewPolicy) Update(c *fiber.Ctx, input requests.ReviewUpdate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameReviewEdit)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Delete policy ensures that the user has the right permissions for delete a review.
func (policy *ReviewPolicy) Delete(c *fiber.Ctx, input requests.ReviewDelete) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameCourseDelete)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}
