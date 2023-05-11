package requests

import "github.com/gofiber/fiber/v2"

type DiscussionDelete struct {
	Id string `json:"id" validate:"required"`
}

func (input *DiscussionDelete) FromContext(c *fiber.Ctx) (err error) {
	if id := c.Params("discussion_id"); id != "" {
		input.Id = id
	} else {
		err = c.BodyParser(input)
	}
	return
}
