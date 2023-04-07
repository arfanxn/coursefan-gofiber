package seeders

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

type UserSeeder struct {
	repository *repositories.UserRepository
}

// NewUserSeeder instantiates a new UserSeeder
func NewUserSeeder(repository *repositories.UserRepository) *UserSeeder {
	return &UserSeeder{
		repository: repository,
	}
}

// Run runs the seeder
func (seeder *UserSeeder) Run(c *fiber.Ctx) (err error) {
	var users []*models.User

	users = append(users, &models.User{
		Name:      "Muhammad Arfan",
		Email:     "arf@gm.com",
		Password:  string(errorh.Must(bcrypt.GenerateFromPassword([]byte("111222"), bcrypt.DefaultCost))),
		CreatedAt: time.Now(),
		UpdatedAt: sliceh.Random(
			null.NewTime(time.Now(), true),
			null.NewTime(time.Time{}, false),
		),
	})

	for i := 0; i < 49; i++ {
		user := factories.FakeUser()
		users = append(users, &user)
	}
	_, err = seeder.repository.Insert(c, users...)

	return
}
