package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

func FakeCourse() models.Course {
	return models.Course{
		// Id:, // will be filled in later
		// Name:, // will be filled in later
		// Slug:, // will be filled in later
		// Description:, // will be filled in later
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
