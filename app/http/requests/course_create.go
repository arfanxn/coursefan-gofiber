package requests

import (
	"github.com/gofiber/fiber/v2"
)

type CourseCreate struct {
	Name        string `json:"name" validate:"required,min=2,max50"`
	Description string `json:"description" validate:"required,min=100"`
}

// FromContext fills input from the given context
func (input *CourseCreate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	return
}
