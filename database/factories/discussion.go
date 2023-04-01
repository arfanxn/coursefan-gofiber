package factories

import (
	"math/rand"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakeDiscussion() models.Discussion {
	return models.Discussion{
		// Id:, // will be filled in later
		// DiscussableType:, // will be filled in later
		// DiscussableId:, // will be filled in later
		// DiscusserId:, // will be filled in later
		// Discusser:, // will be filled in later
		// DiscussionRepliedId:, // will be filled in later
		// DiscussionReplied:, // will be filled in later
		Title: faker.Word(),
		Body: sliceh.Random(
			null.NewString(faker.Sentence(), true),
			null.NewString("", false),
		),
		Upvote:    rand.Intn(100),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
