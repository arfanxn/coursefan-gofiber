package factories

import (
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

// FakeUserSetting returns fake data of models.UserSetting
func FakeUserSetting() models.UserSetting {
	return models.UserSetting{
		// Id: will be filled in later
		// UserId: will be filled in later
		// User: will be filled in later
		Key:       sliceh.Random(enums.UserSettingKeys()...),
		Value:     strconv.FormatBool(sliceh.Random(true, false)),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	}
}
