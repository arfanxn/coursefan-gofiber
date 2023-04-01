package factories

import (
	"math/rand"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakePart() models.Part {
	return models.Part{
		// Id:, // will be filled in later
		// PartableType:, // will be filled in later
		// PartableId:, // will be filled in later
		Part: rand.Intn(10),
		Name: faker.Word(),
		Description:sliceh.Random(
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
