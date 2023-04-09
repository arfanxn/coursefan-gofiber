package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
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
		Column: "id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("lecture_part_id")},
	}, requests.QueryFilter{
		Column: "course_id", Operator: enums.QueryFilterOperatorEquals, Values: []any{c.Params("course_id")},
	})

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
