package requests

import "github.com/gofiber/fiber/v2"

type LecturePartDelete struct {
	Id       string `json:"id" validate:"required,uuid"`
	CourseId string `json:"course_id" validate:"required,uuid"`
}

func (input *LecturePartDelete) FromContext(c *fiber.Ctx) (err error) {
	if courseId := c.Params("course_id"); courseId != "" {
		input.Id = c.Params("lecture_part_id")
		input.CourseId = courseId
	} else {
		err = c.BodyParser(input)
	}
	return
}
