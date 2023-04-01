package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

func FakeToken() models.Token {
	return models.Token{
		// Id:, // will be filled in later
		// TokenableType:, // will be filled in later
		// TokenableId:, // will be filled in later
		Type: enums.TokenTypeResetPassword,
		Body: "11122",
		UsedAt: sliceh.Random(
			null.NewTime(time.Now().Add(-(time.Hour/2)), true),
			null.NewTime(time.Time{}, false),
		),
		ExpiredAt: sliceh.Random(
			null.NewTime(time.Now().Add(time.Hour/2), true),
			null.NewTime(time.Time{}, false),
		),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
