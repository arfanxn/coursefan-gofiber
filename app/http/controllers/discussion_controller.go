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

type DiscussionController struct {
	policy  *policies.DiscussionPolicy
	service *services.DiscussionService
}

func NewDiscussionController(
	policy *policies.DiscussionPolicy,
	service *services.DiscussionService,
) *DiscussionController {
	return &DiscussionController{
		policy:  policy,
		service: service,
	}
}

// All By Lecture
func (controller *DiscussionController) AllByLecture(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.Discussion{}.TableName() + ".discussable_id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{c.Params("lecture_id")},
	})
	err = controller.policy.AllByLecture(c, input)
	if err != nil {
		return
	}
	pagination, err := controller.service.AllByLecture(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve lecture discussions",
		Pagination: pagination,
	})
}

// Find
func (controller *DiscussionController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: models.Discussion{}.TableName() + ".id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("discussion_id")},
	})
	err = controller.policy.Find(c, input)
	if err != nil {
		return
	}
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve discussion",
		Data:    data,
	})
}

// Create
func (controller *DiscussionController) Create(c *fiber.Ctx) (err error) {
	input := requests.DiscussionCreate{}
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
		Message: "Successfully create discussion",
		Data:    data,
	})
}

// Create By Lecture
func (controller *DiscussionController) CreateByLecture(c *fiber.Ctx) (err error) {
	input := requests.DiscussionCreate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.DiscussableType = reflecth.GetTypeName(models.Lecture{})
	input.DiscussableId = c.Params("lecture_id")
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.CreateByLecture(c, input)
	if err != nil {
		return
	}
	data, err := controller.service.CreateByLecture(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully create discussion",
		Data:    data,
	})
}

// Update
func (controller *DiscussionController) Update(c *fiber.Ctx) (err error) {
	input := requests.DiscussionUpdate{}
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
		Message: "Successfully update discussion",
		Data:    data,
	})
}

// Upvote
func (controller *DiscussionController) Upvote(c *fiber.Ctx) (err error) {
	input := requests.DiscussionId{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.Upvote(c, input)
	if err != nil {
		return
	}
	data, err := controller.service.Upvote(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully update discussion",
		Data:    data,
	})
}

// Delete
func (controller *DiscussionController) Delete(c *fiber.Ctx) (err error) {
	input := requests.DiscussionId{}
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
		Message: "Successfully delete discussion",
	})
}
