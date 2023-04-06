package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type LecturePart struct {
	Id        string    `json:"id"`
	CourseId  string    `json:"course_id"`
	Course    *Course   `json:"course,omitempty"`
	Part      int       `json:"part"`
	Name      string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func (resource *LecturePart) FromModel(model models.LecturePart) {
	resource.Id = model.Id.String()
	resource.CourseId = model.CourseId.String()
	if model.Course != nil {
		course := Course{}
		course.FromModel(*model.Course)
		resource.Course = &course
	}
	resource.Part = model.Part
	resource.Name = model.Name
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
