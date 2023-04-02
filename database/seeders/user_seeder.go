package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
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
	for i := 0; i < 50; i++ {
		user := factories.FakeUser()
		users = append(users, &user)
	}
	_, err = seeder.repository.Insert(c, users...)

	return
}
