package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type ProgressUserSeeder struct {
	repository         *repositories.ProgressUserRepository
	progressRepository *repositories.ProgressRepository
	userRepository     *repositories.UserRepository
}

// NewProgressUserSeeder instantiates a new ProgressUserSeeder
func NewProgressUserSeeder(
	repository *repositories.ProgressUserRepository,
	progressRepository *repositories.ProgressRepository,
	userRepository *repositories.UserRepository,
) *ProgressUserSeeder {
	return &ProgressUserSeeder{
		repository:         repository,
		progressRepository: progressRepository,
		userRepository:     userRepository,
	}
}

// Run runs the seeder
func (seeder *ProgressUserSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	// get all progress
	progresses, err := seeder.progressRepository.All(c)
	if err != nil {
		return
	}
	// get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	var progressUsers []*models.ProgressUser
	for _, progress := range progresses {
		progressUser := factories.FakeProgressUser()
		progressUser.ProgressId = progress.Id
		progressUser.UserId = sliceh.Random(users...).Id
		progressUsers = append(progressUsers, &progressUser)
	}

	_, err = seeder.repository.Insert(c, progressUsers...)
	return
}
