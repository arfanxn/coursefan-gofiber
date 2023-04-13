package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type Lecture struct {
	Id            string       `json:"id"`
	LecturePartId string       `json:"lecture_part_id"`
	LecturePart   *LecturePart `json:"lecture_part,omitempty"`
	Name          string       `json:"name"`
	Order         int          `json:"order"`
	VideoUrl      string       `json:"video_url"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     null.Time    `json:"updated_at"`
}

func (resource *Lecture) FromModel(model models.Lecture) {
	resource.Id = model.Id.String()
	resource.LecturePartId = model.LecturePartId.String()
	if model.LecturePart != nil {
		lecturePartRes := LecturePart{}
		lecturePartRes.FromModel(*model.LecturePart)
		resource.LecturePart = &lecturePartRes
	}
	resource.Name = model.Name
	resource.Order = model.Order
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
