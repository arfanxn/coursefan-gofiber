package requests

import "github.com/gofiber/fiber/v2"

type CourseDelete struct {
	Id string `json:"id" validate:"required"`
}

func (input *CourseDelete) FromContext(c *fiber.Ctx) (err error) {
	if id := c.Params("id"); id != "" {
		input.Id = id
	} else {
		err = c.BodyParser(input)
	}
	return
}
