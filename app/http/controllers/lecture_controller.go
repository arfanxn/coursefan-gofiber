package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/validatorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/policies"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type LectureController struct {
	policy  *policies.LecturePolicy
	service *services.LectureService
}

func NewLectureController(
	policy *policies.LecturePolicy,
	service *services.LectureService,
) *LectureController {
	return &LectureController{
		policy:  policy,
		service: service,
	}
}

func (controller *LectureController) AllByLecturePart(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: "lectures.lecture_part_id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("lecture_part_id")},
	})
	err = controller.policy.AllByLecturePart(c, input)
	if err != nil {
		return err
	}
	pagination, err := controller.service.AllByLecturePart(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve lectures",
		Pagination: pagination,
	})
}

// Find
func (controller *LectureController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: "lectures.lecture_part_id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("lecture_part_id")},
	}, requests.QueryFilter{
		Column: "lectures.id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("lecture_id")},
	})
	err = controller.policy.Find(c, input)
	if err != nil {
		return err
	}
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}

	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve lecture",
		Data:    data,
	})
}

// Create
func (controller *LectureController) Create(c *fiber.Ctx) (err error) {
	input := requests.LectureCreate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.Create(c, input)
	if err != nil {
		return err
	}
	data, err := controller.service.Create(c, input)
	if err != nil {
		return err
	}

	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully create lecture",
		Data:    data,
	})
}

// Update
func (controller *LectureController) Update(c *fiber.Ctx) (err error) {
	input := requests.LectureUpdate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.Update(c, input)
	if err != nil {
		return err
	}
	data, err := controller.service.Update(c, input)
	if err != nil {
		return err
	}

	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully update lecture",
		Data:    data,
	})
}

// Delete
func (controller *LectureController) Delete(c *fiber.Ctx) (err error) {
	input := requests.LectureDelete{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	err = controller.policy.Delete(c, input)
	if err != nil {
		return err
	}
	err = controller.service.Delete(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully delete lecture",
	})
}
