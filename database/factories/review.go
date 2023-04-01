package factories

import (
	"math/rand"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakeReview() models.Review {
	return models.Review{
		// Id:, // will be filled in later
		// ReviewableType:, // will be filled in later
		// ReviewableId:, // will be filled in later
		// ReviewerId:, // will be filled in later
		// Reviewer:, // will be filled in later
		Rate: rand.Intn(5),
		Title: sliceh.Random(
			null.NewString(faker.Word(), true),
			null.NewString("", false),
		),
		Body: sliceh.Random(
			null.NewString(faker.Sentence(), true),
			null.NewString("", false),
		),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
