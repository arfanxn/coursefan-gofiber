package requests

import (
	"github.com/gofiber/fiber/v2"
)

type UserSettingUpdate struct {
	UserId string `json:"user_id" form:"user_id" validate:"required,uuid"`
	Key    string `json:"key" form:"key" validate:"required,oneof=user_public_profile user_public_course notification_accept_promotion"`
	Value  string `json:"value" form:"value" validate:"required"`
}

// FromContext fills input from the given context
func (input *UserSettingUpdate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	if settingKey := c.Params("user_setting_key"); settingKey != "" {
		input.Key = settingKey
	}

	return
}
