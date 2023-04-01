package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakeMessage() models.Message {
	return models.Message{
		//Id:, // will be filled in later
		//SenderId:, // will be filled in later
		//Sender:, // will be filled in later
		//ReceiverId:, // will be filled in later
		//Receiver:, // will be filled in later
		Body: faker.Sentence(),
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
