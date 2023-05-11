package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type DiscussionPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewDiscussionPolicy(permissionRepository *repositories.PermissionRepository) *DiscussionPolicy {
	return &DiscussionPolicy{permissionRepository: permissionRepository}
}

// AllByLecture policy ensures that the user has the right permissions for access a course discussions.
func (policy *DiscussionPolicy) AllByLecture(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLectureDiscussionView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Find policy ensures that the user has the right permissions for access a discussion.
func (policy *DiscussionPolicy) Find(c *fiber.Ctx, input requests.Query) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameDiscussionView)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Create policy ensures that the user has the right permissions for create a discussion.
func (policy *DiscussionPolicy) Create(c *fiber.Ctx, input requests.DiscussionCreate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameDiscussionCreate)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// CreateByLecture policy ensures that the user has the right permissions for create a discussion by the given course.
func (policy *DiscussionPolicy) CreateByLecture(c *fiber.Ctx, input requests.DiscussionCreate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameLectureDiscussionCreate)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Update policy ensures that the user has the right permissions for update a discussion.
func (policy *DiscussionPolicy) Update(c *fiber.Ctx, input requests.DiscussionUpdate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameDiscussionEdit)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Upvote policy ensures that the user has the right permissions for upvote a discussion.
func (policy *DiscussionPolicy) Upvote(c *fiber.Ctx, input requests.DiscussionId) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameDiscussionUpvote)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Delete policy ensures that the user has the right permissions for delete a discussion.
func (policy *DiscussionPolicy) Delete(c *fiber.Ctx, input requests.DiscussionId) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameDiscussionDelete)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}
