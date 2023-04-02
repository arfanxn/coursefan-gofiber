package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

// FakeUserProfile returns fake data of models.UserProfile
func FakeUserProfile() models.UserProfile {
	return models.UserProfile{
		// Id: will be filled in later
		// UserId: will be filled in later
		// User: will be filled in later
		Headline: sliceh.Random(
			null.NewString(faker.Word(), true),
			null.NewString("", false),
		),
		Biography: sliceh.Random(
			null.NewString(faker.Sentence(), true),
			null.NewString("", false),
		),
		Language: sliceh.Random("en", "id"),
		WebsiteUrl: sliceh.Random(
			null.NewString(faker.URL(), true),
			null.NewString("", false),
		),
		FacebookUrl: sliceh.Random(
			null.NewString(faker.URL(), true),
			null.NewString("", false),
		),
		LinkedinUrl: sliceh.Random(
			null.NewString(faker.URL(), true),
			null.NewString("", false),
		),
		TwitterUrl: sliceh.Random(
			null.NewString(faker.URL(), true),
			null.NewString("", false),
		),
		YoutubeUrl: sliceh.Random(
			null.NewString(faker.URL(), true),
			null.NewString("", false),
		),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
