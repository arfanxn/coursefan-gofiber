package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakeNotification() models.Notification {
	return models.Notification{
		// Id:, // will be filled in later
		// SenderId:, // will be filled in later
		// Sender:, // will be filled in later
		// ReceiverId:, // will be filled in later
		// Receiver:, // will be filled in later
		// ObjectType:, // will be filled in later
		// ObjectId:, // will be filled in later
		Title: faker.Word(),
		Body: sliceh.Random(
			null.NewString(faker.Sentence(), true),
			null.NewString("", false),
		),
		// Type:, // will be filled in later
		ReadedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
