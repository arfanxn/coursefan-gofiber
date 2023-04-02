package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type MessageSeeder struct {
	repository     *repositories.MessageRepository
	userRepository *repositories.UserRepository
}

// NewMessageSeeder instantiates a new MessageSeeder
func NewMessageSeeder(
	repository *repositories.MessageRepository,
	userRepository *repositories.UserRepository,
) *MessageSeeder {
	return &MessageSeeder{
		repository:     repository,
		userRepository: userRepository,
	}
}

// Run runs the seeder
func (seeder *MessageSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var messages []*models.Message
	for _, user := range users {
		for i := 0; i < 2; i++ {
			message := factories.FakeMessage()
			message.SenderId = users[0].Id
			message.ReceiverId = user.Id
			messages = append(messages, &message)
		}
	}
	_, err = seeder.repository.Insert(c, messages...)

	return
}
