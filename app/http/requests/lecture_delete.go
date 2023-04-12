package requests

import "github.com/gofiber/fiber/v2"

type LectureDelete struct {
	Id            string `json:"id" validate:"required,uuid"`
	LecturePartId string `json:"course_id" validate:"required,uuid"`
}

func (input *LectureDelete) FromContext(c *fiber.Ctx) (err error) {
	if lectureId := c.Params("lecture_id"); lectureId != "" {
		input.Id = lectureId
		input.LecturePartId = c.Params("lecture_part_id")
	} else {
		err = c.BodyParser(input)
	}
	return
}
