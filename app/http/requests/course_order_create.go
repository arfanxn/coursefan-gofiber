package requests

import (
	"github.com/gofiber/fiber/v2"
)

type CourseOrderCreate struct {
	UserId   string `json:"user_id" form:"user_id" validate:"required,uuid"`
	CourseId string `json:"course_id" form:"course_id" validate:"required,uuid"`
	Bank     string `json:"bank" form:"bank" validate:"required,oneof=bni bca"`
}

// FromContext fills input from the given context
func (input *CourseOrderCreate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	if courseId := c.Params("course_id"); courseId != "" {
		input.CourseId = courseId
	}

	return
}
