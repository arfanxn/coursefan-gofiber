package factories

import (
	"database/sql"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/go-faker/faker/v4"
)

// User returns fake data of models.User
func User() models.User {
	hashedPassword := "$2a$10$1sGm.uAbtb6h9HkZv1/5S.IFesDq7GOJx0gjXAhGltA3hFssCs/kO" // unhashedPassword = 111222
	return models.User{
		Name:      faker.Name(),
		Email:     faker.Email(),
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			sql.NullTime{Time: time.Time{}, Valid: false},
			sql.NullTime{Time: time.Now(), Valid: true},
		),
	}
}
