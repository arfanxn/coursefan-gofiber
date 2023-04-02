package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type UserSettingSeeder struct {
	repository     *repositories.UserSettingRepository
	userRepository *repositories.UserRepository
}

// NewUserSettingSeeder instantiates a new UserSettingSeeder
func NewUserSettingSeeder(
	repository *repositories.UserSettingRepository,
	userRepository *repositories.UserRepository,
) *UserSettingSeeder {
	return &UserSettingSeeder{
		repository:     repository,
		userRepository: userRepository,
	}
}

// Run runs the seeder
func (seeder *UserSettingSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var userSettings []*models.UserSetting
	for _, user := range users {
		userSetting := factories.FakeUserSetting()
		userSetting.UserId = user.Id
		userSettings = append(userSettings, &userSetting)
	}
	_, err = seeder.repository.Insert(c, userSettings...)

	return
}
