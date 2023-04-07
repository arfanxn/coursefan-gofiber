package requests

import (
	"github.com/gofiber/fiber/v2"
)

type CourseUpdate struct {
	Id          string `json:"id" validate:"required,uuid"`
	Name        string `json:"name" validate:"required,min=2,max=50"`
	Description string `json:"description" validate:"required,min=100"`
}

// FromContext fills input from the given context
func (input *CourseUpdate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	if id := c.Params("id"); id != "" {
		input.Id = id
	}
	return
}
