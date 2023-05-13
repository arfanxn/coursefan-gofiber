package factories

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"github.com/iancoleman/strcase"
	"gopkg.in/guregu/null.v4"
)

func FakeCourse() models.Course {
	courseName := faker.Word() + " " + faker.Word() + strconv.FormatInt(rand.Int63n(9999)+int64(1000), 10)
	return models.Course{
		// Id:, // will be filled in later
		Name:        courseName,
		Slug:        strcase.ToKebab(courseName), // will be filled in later
		Description: faker.Sentence(),            // will be filled in later
		Price:       float64(0 + rand.Int63n(200000)),
		CreatedAt:   time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
