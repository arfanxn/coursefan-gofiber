package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/validatorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/policies"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type ReviewController struct {
	policy  *policies.ReviewPolicy
	service *services.ReviewService
}

func NewReviewController(
	policy *policies.ReviewPolicy,
	service *services.ReviewService,
) *ReviewController {
	return &ReviewController{
		policy:  policy,
		service: service,
	}
}

// All By Course
func (controller *ReviewController) AllByCourse(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: models.Review{}.TableName() + ".reviewable_id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("course_id")},
	})
	// No api policy required, always allow anyone to access this resource
	pagination, err := controller.service.AllByCourse(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve course reviews",
		Pagination: pagination,
	})
}

// Find
func (controller *ReviewController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: models.Review{}.TableName() + ".id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("review_id")},
	})
	// No api policy required, always allow anyone to access this resource
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve review",
		Data:    data,
	})
}

// Create
func (controller *ReviewController) Create(c *fiber.Ctx) (err error) {
	input := requests.ReviewCreate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.Create(c, input)
	if err != nil {
		return
	}
	data, err := controller.service.Create(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully create review",
		Data:    data,
	})
}

// Create By Course
func (controller *ReviewController) CreateByCourse(c *fiber.Ctx) (err error) {
	input := requests.ReviewCreate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.ReviewableType = reflecth.GetTypeName(models.Course{})
	input.ReviewableId = c.Params("course_id")
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.CreateByCourse(c, input)
	if err != nil {
		return
	}
	data, err := controller.service.CreateByCourse(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully create review",
		Data:    data,
	})
}

// Update
func (controller *ReviewController) Update(c *fiber.Ctx) (err error) {
	input := requests.ReviewUpdate{}
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
		Message: "Successfully update review",
		Data:    data,
	})
}

// Delete
func (controller *ReviewController) Delete(c *fiber.Ctx) (err error) {
	input := requests.ReviewDelete{}
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
		Message: "Successfully delete review",
	})
}
