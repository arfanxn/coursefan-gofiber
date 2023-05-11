package requests

import "github.com/gofiber/fiber/v2"

type DiscussionId struct {
	Id string `json:"id" validate:"required"`
}

func (input *DiscussionId) FromContext(c *fiber.Ctx) (err error) {
	if id := c.Params("discussion_id"); id != "" {
		input.Id = id
	} else {
		err = c.BodyParser(input)
	}
	return
}
