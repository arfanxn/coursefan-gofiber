package requests

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

type NotificationCreate struct {
	SenderId   string      `json:"sender_id" form:"sender_id" validate:"required,uuid"`
	ReceiverId string      `json:"receiver_id" form:"receiver_id" validate:"required,uuid"`
	ObjectType null.String `json:"object_type" form:"object_type"`
	ObjectId   null.String `json:"object_id" form:"object_id"`
	Title      string      `json:"title" form:"title" validate:"required"`
	Body       null.String `json:"body" form:"body"`
	Type       null.String `json:"type" form:"type"`
}

// FromContext fills input from the given context
func (input *NotificationCreate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	return
}
