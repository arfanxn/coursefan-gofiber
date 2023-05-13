package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/boolh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/funch"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/validatorh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/policies"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CourseOrderController struct {
	policy  *policies.CourseOrderPolicy
	service *services.CourseOrderService
}

func NewCourseOrderController(
	policy *policies.CourseOrderPolicy,
	service *services.CourseOrderService,
) *CourseOrderController {
	return &CourseOrderController{
		policy:  policy,
		service: service,
	}
}

// All By Auth User
func (controller *CourseOrderController) AllByAuthUser(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.CourseOrder{}.TableName() + ".user_id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{ctxh.MustGetUser(c).Id},
	})
	// No api policy required, only the authenticated user can access their own course orders
	pagination, err := controller.service.All(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve course orders",
		Pagination: pagination,
	})
}

// CreateByAuthUser
func (controller *CourseOrderController) CreateByAuthUser(c *fiber.Ctx) (err error) {
	input := requests.CourseOrderCreate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.UserId = ctxh.MustGetUser(c).Id.String()
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// No api policy required, only the authenticated user can create a course order
	data, err := controller.service.Create(c, input)
	if err != nil {
		return err
	}

	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully create course order",
		Data: map[string]any{
			"charge":       data.CoreapiChargeResponse,
			"course_order": data.ResourceCourseOrder,
		},
	})
}

// Find
func (controller *CourseOrderController) Find(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.CourseOrder{}.TableName() + ".id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{c.Params("course_order_id")},
	})
	// No api policy required, only the authenticated user can access course order
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve course order",
		Data:    data,
	})
}

// UpdateByMidtransNotification
func (controller *CourseOrderController) UpdateByMidtransNotification(c *fiber.Ctx) (err error) {
	input := requests.MidtransNotification{}
	err = input.FromContext(c)
	if err != nil {
		return
	}

	// No api policy required, only the authenticated user can access course order

	go func() { // Do others on the goroutine thread so that we don't have to wait for the whole process to finish, we can just immediately return to the client side
		var err error = nil
		funch.Recursive(func(repeat *bool) {
			_, err = controller.service.UpdateByMidtransNotification(c, input)
			if err != nil {
				repeat = boolh.ToPointer(true)
				return
			}
			repeat = boolh.ToPointer(false)
		}, 5)
		if err != nil { // if error occurs after recursion call then log the error
			logrus.Error(err)
		}
	}()

	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully update course order",
	})
}
