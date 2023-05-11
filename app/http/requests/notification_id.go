package requests

import (
	"github.com/gofiber/fiber/v2"
)

type NotificationId struct {
	Id string `json:"id" form:"id" validate:"required,uuid"`
}

// FromContext fills input from the given context
func (input *NotificationId) FromContext(c *fiber.Ctx) (err error) {
	if id := c.Params("notification_id"); id != "" {
		input.Id = id
	} else {
		err = c.BodyParser(input)
	}
	return
}
