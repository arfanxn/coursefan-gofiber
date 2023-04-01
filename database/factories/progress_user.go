package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

func FakeProgressUser() models.ProgressUser {
	return models.ProgressUser{
		// Id:, will be filled in later
		// ProgressId:, will be filled in later
		// Progress:, will be filled in later
		// UserId:, will be filled in later
		// User:, will be filled in later
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
