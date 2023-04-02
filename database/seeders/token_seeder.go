package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type TokenSeeder struct {
	repository     *repositories.TokenRepository
	userRepository *repositories.UserRepository
}

// NewTokenSeeder instantiates a new TokenSeeder
func NewTokenSeeder(
	repository *repositories.TokenRepository,
	userRepository *repositories.UserRepository,
) *TokenSeeder {
	return &TokenSeeder{
		repository:     repository,
		userRepository: userRepository,
	}
}

// Run runs the seeder
func (seeder *TokenSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var tokens []*models.Token
	for _, user := range users {
		token := factories.FakeToken()
		token.TokenableType = reflecth.GetTypeName(user)
		token.TokenableId = user.Id
		tokens = append(tokens, &token)
	}
	_, err = seeder.repository.Insert(c, tokens...)

	return
}
