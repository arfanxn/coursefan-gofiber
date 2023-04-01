package factories

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
)

func FakePermission() models.Permission {
	return models.Permission{
		// Id:, // will be filled in later
		Name: sliceh.Random(enums.PermissionNames()...),
	}
}
