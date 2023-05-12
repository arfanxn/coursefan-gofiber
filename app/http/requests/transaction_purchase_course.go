package requests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TransactionPurchaseCourse struct {
	TransactionableId uuid.UUID `json:"transactionable_id" form:"transactionable_id" validate:"required,uuid"`
	SenderId          uuid.UUID `json:"sender_id" form:"sender_id" validate:"required,uuid"`
}

// FromContext fills input from the given context
func (input *TransactionPurchaseCourse) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	return
}
