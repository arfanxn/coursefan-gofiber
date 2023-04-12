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

type LecturePartController struct {
	policy  *policies.LecturePartPolicy
	service *services.LecturePartService
}

func NewLecturePartController(
	policy *policies.LecturePartPolicy,
	service *services.LecturePartService,
) *LecturePartController {
	return &LecturePartController{
		policy:  policy,
		service: service,
	}
}

func (controller *LecturePartController) AllByCourse(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: "lecture_parts.course_id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("course_id")},
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
		Message:    "Successfully retrieve lecture parts",
		Pagination: pagination,
	})
}

// Find
func (controller *LecturePartController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column: "lecture_parts.id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("lecture_part_id")},
	}, requests.QueryFilter{
		Column: "lecture_parts.course_id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("course_id")},
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
		Message: "Successfully retrieve lecture part",
		Data:    data,
	})
}

// Create
func (controller *LecturePartController) Create(c *fiber.Ctx) (err error) {
	input := requests.LecturePartCreate{}
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
		Code:    fiber.StatusOK,
		Message: "Successfully create lecture part",
		Data:    data,
	})
}

// Update
func (controller *LecturePartController) Update(c *fiber.Ctx) (err error) {
	input := requests.LecturePartUpdate{}
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
		Message: "Successfully update lecture part",
		Data:    data,
	})
}

// Delete
func (controller *LecturePartController) Delete(c *fiber.Ctx) (err error) {
	input := requests.LecturePartDelete{}
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
		Message: "Successfully delete lecture part",
	})
}
