package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakeLecturePart() models.LecturePart {
	return models.LecturePart{
		// Id:, // will be filled in later
		// CourseId:, // will be filled in later
		// Course:, // will be filled in later
		// Part:, // will be filled in later
		Name:      faker.Word(),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
