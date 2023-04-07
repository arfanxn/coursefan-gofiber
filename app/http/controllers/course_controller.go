package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type CourseController struct {
	service *services.CourseService
}

// NewCourseController instantiates a new CourseController
func NewCourseController(service *services.CourseService) *CourseController {
	return &CourseController{service: service}
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
