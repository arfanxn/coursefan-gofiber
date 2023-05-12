package requests

import (
	"github.com/gofiber/fiber/v2"
)

type TransactionPurchaseCourse struct {
	TransactionableId string `json:"transactionable_id" form:"transactionable_id" validate:"required,uuid"`
	SenderId          string `json:"sender_id" form:"sender_id" validate:"required,uuid"`
}

// FromContext fills input from the given context
func (input *TransactionPurchaseCourse) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	if courseId := c.Params("course_id"); courseId != "" {
		input.TransactionableId = courseId
	}

	return
}
