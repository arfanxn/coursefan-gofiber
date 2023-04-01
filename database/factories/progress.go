package factories

import (
	"math/rand"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

func FakeProgress() models.Progress {
	return models.Progress{
		// Id:, // will be filled in later
		// ProgressableType:, // will be filled in later
		// ProgressableId:, // will be filled in later
		Percentage: rand.Intn(100),
		CreatedAt:  time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
