package factories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
	"gopkg.in/guregu/null.v4"
)

func FakeMedia() models.Media {
	return models.Media{
		// Id:, // will be filled in later
		// ModelType:, // will be filled in later
		// ModelId:, // will be filled in later
		CollectionName: sliceh.Random(
			null.NewString(sliceh.Random(enums.MediaCollectionNames()...), true),
			null.NewString("", false),
		),
		Name: sliceh.Random(
			null.NewString(faker.Word(), true),
			null.NewString("", false),
		),
		// FileName:, // will be filled in later
		// MimeType:, // will be filled in later
		// Disk:, // will be filled in later
		// ConversionDisk:, // will be filled in later
		// Size:, // will be filled in later
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
