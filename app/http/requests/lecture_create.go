package requests

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type LectureCreate struct {
	LecturePartId string `json:"course_id" validate:"required,uuid"`
	int           `json:""  validate:"required,number,min=0"`
	Name          string                `json:"name" validate:"required"`
	Order         int                   `json:"order" validate:"required,min=1"`
	Video         *multipart.FileHeader `json:"video" fhlidate:"required,max=500"`
}

func (input *LectureCreate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	if err != nil {
		return
	}
	input.LecturePartId = c.Params("lecture_part_id")
	lectureVideo, err := c.FormFile("video")
	if err != nil {
		return
	}
	input.Video = lectureVideo
	return
}
