package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"github.com/iancoleman/strcase"
	"gopkg.in/guregu/null.v4"
)

func FakeCourse() models.Course {
	courseName := faker.Word() + " " + faker.Word()
	return models.Course{
		// Id:, // will be filled in later
		Name:        courseName,
		Slug:        strcase.ToKebab(courseName), // will be filled in later
		Description: faker.Sentence(),            // will be filled in later
		CreatedAt:   time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
