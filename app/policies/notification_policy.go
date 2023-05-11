package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
)

type NotificationPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewNotificationPolicy(permissionRepository *repositories.PermissionRepository) *NotificationPolicy {
	return &NotificationPolicy{permissionRepository: permissionRepository}
}

/*
! Disabled due the policies are written inline in the service class of notification
// Create policy ensures that the user has the right permissions for create a discussion.
func (policy *NotificationPolicy) Create(c *fiber.Ctx, input requests.NotificationCreate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameNotificationCreate)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Update policy ensures that the user has the right permissions for update a discussion.
func (policy *NotificationPolicy) Update(c *fiber.Ctx, input requests.NotificationUpdate) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameNotificationEdit)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// MarkReadOrUnread policy ensures that the user has the right permissions for marking a notification as read or unreaded.
func (policy *NotificationPolicy) MarkReadOrUnread(c *fiber.Ctx, input requests.NotificationId) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameNotificationMarkReadOrUnread)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}

// Delete policy ensures that the user has the right permissions for delete a discussion.
func (policy *NotificationPolicy) Delete(c *fiber.Ctx, input requests.NotificationId) (err error) {
	_, err = policy.permissionRepository.FindByNameAndCUR(c, enums.PermissionNameNotificationDelete)
	if errorh.IsGormErrRecordNotFound(err) {
		return fiber.ErrForbidden
	} else if err != nil {
		return
	}
	return nil
}
*/
