package requests

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

type ReviewUpdate struct {
	Id    string      `json:"id" validate:"required,uuid"`
	Rate  int         `json:"rate"  validate:"required,number"`
	Title null.String `json:"title"`
	Body  null.String `json:"body"`
}

// FromContext fills input from the given context
func (input *ReviewUpdate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	if id := c.Params("review_id"); id != "" {
		input.Id = id
	}

	return
}
