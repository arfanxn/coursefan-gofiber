package requests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserSettingUpdate struct {
	UserId uuid.UUID `json:"user_id" validate:"required,uuid"`
	Key    string    `json:"key" validate:"required"`
	Value  string    `json:"value" validate:"required"`
}

// FromContext fills input from the given context
func (input *UserSettingUpdate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	if settingKey := c.Params("user_setting_key"); settingKey != "" {
		input.Key = settingKey
	}

	return
}
