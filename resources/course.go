package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type Course struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Slug         string        `json:"slug"`
	Description  string        `json:"description"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    null.Time     `json:"updated_at"`
	LectureParts []LecturePart `json:"lecture_parts,omitempty"`
}

// FromModel fills the resource from the given model
func (resource *Course) FromModel(model models.Course) {
	resource.Id = model.Id.String()
	resource.Name = model.Name
	resource.Slug = model.Slug
	resource.Description = model.Description
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
	for _, lecturePartMdl := range model.LectureParts {
		lecturePartRes := LecturePart{}
		lecturePartRes.FromModel(lecturePartMdl)
		resource.LectureParts = append(resource.LectureParts, lecturePartRes)
	}
}
