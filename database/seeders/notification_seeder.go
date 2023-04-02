package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type NotificationSeeder struct {
	repository     *repositories.NotificationRepository
	userRepository *repositories.UserRepository
}

// NewNotificationSeeder instantiates a new NotificationSeeder
func NewNotificationSeeder(
	repository *repositories.NotificationRepository,
	userRepository *repositories.UserRepository,
) *NotificationSeeder {
	return &NotificationSeeder{
		repository:     repository,
		userRepository: userRepository,
	}
}

// Run runs the seeder
func (seeder *NotificationSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var notifications []*models.Notification
	for _, user := range users {
		for i := 0; i < 2; i++ {
			notification := factories.FakeNotification()
			notification.SenderId = users[0].Id
			notification.ReceiverId = user.Id
			notifications = append(notifications, &notification)
		}
	}
	_, err = seeder.repository.Insert(c, notifications...)

	return
}
