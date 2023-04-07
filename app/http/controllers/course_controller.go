package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/validatorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/policies"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type CourseController struct {
	policy  *policies.CoursePolicy
	service *services.CourseService
}

// NewCourseController instantiates a new CourseController
func NewCourseController(
	policy *policies.CoursePolicy,
	service *services.CourseService,
) *CourseController {
	return &CourseController{
		policy:  policy,
		service: service,
	}
}

// All
func (controller *CourseController) All(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	pagination, err := controller.service.All(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve courses",
		Pagination: pagination,
	})
}

// Find
func (controller *CourseController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.Filters = append(input.Filters, requests.QueryFilter{
		Column: "id", Operator: "==", Values: []any{c.Params("id")},
	})

	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}

	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve course",
		Data:    data,
	})
}

// Create
func (controller *CourseController) Create(c *fiber.Ctx) (err error) {
	input := requests.CourseCreate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	data, err := controller.service.Create(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully create course",
		Data:    data,
	})
}

// Update
func (controller *CourseController) Update(c *fiber.Ctx) (err error) {
	input := requests.CourseUpdate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.Update(c, input)
	if err != nil {
		return
	}
	data, err := controller.service.Update(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully update course",
		Data:    data,
	})
}

// Delete
func (controller *CourseController) Delete(c *fiber.Ctx) (err error) {
	input := requests.CourseDelete{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.Delete(c, input)
	if err != nil {
		return
	}
	err = controller.service.Delete(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully delete course",
	})
}
