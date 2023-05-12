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

type UserSettingController struct {
	policy  *policies.UserSettingPolicy
	service *services.UserSettingService
}

func NewUserSettingController(
	policy *policies.UserSettingPolicy,
	service *services.UserSettingService,
) *UserSettingController {
	return &UserSettingController{
		policy:  policy,
		service: service,
	}
}

// AllByAuthUser
func (controller *UserSettingController) AllByAuthUser(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.UserSetting{}.TableName() + ".user_id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{ctxh.MustGetUser(c).Id.String()},
	})
	// No api policy required, only the authenticated user can access their own user settings
	pagination, err := controller.service.All(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:       fiber.StatusOK,
		Message:    "Successfully retrieve user settings",
		Pagination: pagination,
	})
}

// UpdateByAuthUser updates user setting by the given user_setting_key and the current signed in  user
func (controller *UserSettingController) UpdateByAuthUser(c *fiber.Ctx) (err error) {
	input := requests.UserSettingUpdate{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.UserId = ctxh.MustGetUser(c).Id.String() // assign the current logged in user id
	if err := validatorh.ValidateStruct(input, ctxh.GetAcceptLang(c)); err != nil {
		return err
	}
	// No api policy required, only the authenticated user can update their own user settings
	data, err := controller.service.Update(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully update user setting",
		Data:    data,
	})
}
