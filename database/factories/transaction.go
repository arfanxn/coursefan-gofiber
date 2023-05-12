package factories

import (
	"math/rand"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

func FakeTransaction() models.Transaction {
	return models.Transaction{
		// Id:, // will be filled in later
		// TransactionableType:, // will be filled in later
		// TransactionableId:, // will be filled in later
		// SenderId:, // will be filled in later
		// Sender:, // will be filled in later
		// ReceiverId:, // will be filled in later
		// Receiver:, // will be filled in later
		Amount:   float64(1000 + rand.Int63n(9999)),
		Rate:     float64(10 + rand.Int63n(100)),
		Discount: float64(10 + rand.Int63n(100)),
		Total:    float64(1000 + rand.Int63n(11000)),
		CancelledAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
		ChargebackedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
		ExpiredAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
		FailedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
		RefundedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
		SettledAt: sliceh.Random(
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
