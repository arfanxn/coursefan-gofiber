package requests

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

type ReviewCreate struct {
	ReviewableType string      `json:"reviewable_type" form:"reviewable_type" validate:"oneof=Course User Instructor"`
	ReviewableId   string      `json:"reviewable_id" form:"reviewable_id"  validate:"required,uuid"`
	ReviewerId     string      `json:"reviewer_id" form:"reviewer_id"  validate:"required,uuid"`
	Rate           int         `json:"rate" validate:"required,number"`
	Title          null.String `json:"title"`
	Body           null.String `json:"body"`
}

// FromContext fills input from the given context
func (input *ReviewCreate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	return
}
