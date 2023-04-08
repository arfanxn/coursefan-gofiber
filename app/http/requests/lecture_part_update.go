package requests

import "github.com/gofiber/fiber/v2"

type LecturePartUpdate struct {
	Id       string `json:"id" validate:"required,uuid"`
	CourseId string `json:"course_id" validate:"required,uuid"`
	Part     int    `json:"part"  validate:"required,number,min=0"`
	Name     string `json:"name" validate:"required"`
}

func (input *LecturePartUpdate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	if err != nil {
		return
	}
	input.Id = c.Params("lecture_part_id")
	input.CourseId = c.Params("course_id")
	return
}
