package factories

import (
	"math/rand"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakeLecture() models.Lecture {
	return models.Lecture{
		// Id:, // will be filled in later
		// LecturePartId:, // will be filled in later
		// LecturePart:, // will be filled in later
		Name:      faker.Word(),
		Order:     rand.Intn(10),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
