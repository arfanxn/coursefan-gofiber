package factories

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
)

func FakeRole() models.Role {
	return models.Role{
		// Id:, // will be filled in later
		Name: sliceh.Random(enums.RoleNames()...),
	}
}
