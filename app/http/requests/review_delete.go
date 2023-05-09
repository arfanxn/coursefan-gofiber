package requests

import "github.com/gofiber/fiber/v2"

type ReviewDelete struct {
	Id string `json:"id" validate:"required"`
}

func (input *ReviewDelete) FromContext(c *fiber.Ctx) (err error) {
	if id := c.Params("course_id"); id != "" {
		input.Id = id
	} else {
		err = c.BodyParser(input)
	}
	return
}
