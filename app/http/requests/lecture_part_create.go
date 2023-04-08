package requests

import "github.com/gofiber/fiber/v2"

type LecturePartCreate struct {
	CourseId string `json:"course_id" validate:"required,uuid"`
	Part     int    `json:"part"  validate:"required,number,min=0"`
	Name     string `json:"name" validate:"required"`
}

func (input *LecturePartCreate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	if err != nil {
		return
	}
	input.CourseId = c.Params("course_id")
	return
}
