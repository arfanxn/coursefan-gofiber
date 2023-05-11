package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/validatorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/policies"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	policy  *policies.NotificationPolicy
	service *services.NotificationService
}

func NewNotificationController(
	policy *policies.NotificationPolicy,
	service *services.NotificationService,
) *NotificationController {
	return &NotificationController{
		policy:  policy,
		service: service,
	}
}

// All By Lecture
func (controller *NotificationController) AllByAuthUser(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.Notification{}.TableName() + ".receiver_id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{ctxh.MustGetUser(c).Id},
	})
	// No api policy required, always allow anyone to access this resource
	pagination, err := controller.service.AllByAuthUser(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve notifications",
		Pagination: pagination,
	})
}

// Find
func (controller *NotificationController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: models.Notification{}.TableName() + ".id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("notification_id")},
	})
	// * The api policy is written in the service below
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve notification",
		Data:    data,
	})
}

// Create
func (controller *NotificationController) CreateByAuthUser(c *fiber.Ctx) (err error) {
	input := requests.NotificationCreate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.SenderId = ctxh.MustGetUser(c).Id.String() // fills sender with current auth user id
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// No api policy required, always allow anyone to access this resource
	data, err := controller.service.Create(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully create notification",
		Data:    data,
	})
}

// Update
func (controller *NotificationController) Update(c *fiber.Ctx) (err error) {
	input := requests.NotificationUpdate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// * The api policy is written in the service below
	data, err := controller.service.Update(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully update notification",
		Data:    data,
	})
}

// MarkRead marks the notification as read
func (controller *NotificationController) MarkRead(c *fiber.Ctx) (err error) {
	input := requests.NotificationId{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// * The api policy is written in the service below
	data, err := controller.service.Mark(c, input, true)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully mark notification as read",
		Data:    data,
	})
}

// MarkUnread marks the notification as unread
func (controller *NotificationController) MarkUnread(c *fiber.Ctx) (err error) {
	input := requests.NotificationId{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// * The api policy is written in the service below
	data, err := controller.service.Mark(c, input, false)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully mark notification as unread",
		Data:    data,
	})
}

// Delete
func (controller *NotificationController) Delete(c *fiber.Ctx) (err error) {
	input := requests.NotificationId{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// * The api policy is written in the service below
	err = controller.service.Delete(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully delete notification",
	})
}
