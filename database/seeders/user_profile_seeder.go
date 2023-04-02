package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type UserProfileSeeder struct {
	repository     *repositories.UserProfileRepository
	userRepository *repositories.UserRepository
}

// NewUserProfileSeeder instantiates a new UserProfileSeeder
func NewUserProfileSeeder(
	repository *repositories.UserProfileRepository,
	userRepository *repositories.UserRepository,
) *UserProfileSeeder {
	return &UserProfileSeeder{
		repository:     repository,
		userRepository: userRepository,
	}
}

// Run runs the seeder
func (seeder *UserProfileSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var userProfiles []*models.UserProfile
	for _, user := range users {
		userProfile := factories.FakeUserProfile()
		userProfile.UserId = user.Id
		userProfiles = append(userProfiles, &userProfile)
	}
	_, err = seeder.repository.Insert(c, userProfiles...)

	return
}
