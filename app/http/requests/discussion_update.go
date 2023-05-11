package requests

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

type DiscussionUpdate struct {
	Id    string      `json:"id" validate:"required,uuid"`
	Title string      `json:"title"`
	Body  null.String `json:"body"`
}

// FromContext fills input from the given context
func (input *DiscussionUpdate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	if id := c.Params("discussion_id"); id != "" {
		input.Id = id
	}

	return
}
