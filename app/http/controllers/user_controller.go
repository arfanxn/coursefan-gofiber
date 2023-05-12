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

type UserController struct {
	policy  *policies.UserPolicy
	service *services.UserService
}

func NewUserController(
	policy *policies.UserPolicy,
	service *services.UserService,
) *UserController {
	return &UserController{
		policy:  policy,
		service: service,
	}
}

// All By Course
func (controller *UserController) AllByCourse(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.Course{}.TableName() + ".id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{c.Params("course_id")},
	})
	err = controller.policy.AllByCourse(c, input)
	if err != nil {
		return
	}
	pagination, err := controller.service.AllByCourse(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve course users",
		Pagination: pagination,
	})
}

// Self returns the current logged in user with its associated data.
func (controller *UserController) Self(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.User{}.TableName() + ".id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{ctxh.MustGetUser(c).Id},
	})
	// No api policy required, always allow anyone to access this resource
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve user",
		Data:    data,
	})
}

// Find
func (controller *UserController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.User{}.TableName() + ".id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{c.Params("user_id")},
	})
	// No api policy required, always allow anyone to access this resource
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve user",
		Data:    data,
	})
}

// Update
func (controller *UserController) SelfUpdate(c *fiber.Ctx) (err error) {
	input := requests.UserUpdate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.Id = ctxh.MustGetUser(c).Id.String() // assign the auth user id
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// No api policy required, always allow anyone to access this resource
	data, err := controller.service.Update(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully update user",
		Data:    data,
	})
}
