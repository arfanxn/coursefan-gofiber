package factories

import (
	"math/rand"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

func FakeWallet() models.Wallet {
	return models.Wallet{
		// Id:. // will be filled in later
		// OwnerId:. // will be filled in later
		// Owner:. // will be filled in later
		Balance:   1000 + rand.Int63n(9999),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
