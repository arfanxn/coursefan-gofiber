package seeders

import (
	"math/rand"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DiscussionSeeder struct {
	repository        *repositories.DiscussionRepository
	lectureRepository *repositories.LectureRepository
	userRepository    *repositories.UserRepository
}

// NewDiscussionSeeder instantiates a new DiscussionSeeder
func NewDiscussionSeeder(
	repository *repositories.DiscussionRepository,
	lectureRepository *repositories.LectureRepository,
	userRepository *repositories.UserRepository,
) *DiscussionSeeder {
	return &DiscussionSeeder{
		repository:        repository,
		lectureRepository: lectureRepository,
		userRepository:    userRepository,
	}
}

// Run runs the seeder
func (seeder *DiscussionSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	// Get all lecture
	lectures, err := seeder.lectureRepository.All(c)
	if err != nil {
		return
	}
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var discussions []*models.Discussion
	for _, lecture := range lectures {
		syncronizer.WG().Add(1)
		go func(lecture models.Lecture) {
			defer syncronizer.WG().Done()
			for i := 0; i < rand.Intn(20); i++ {
				discussion := factories.FakeDiscussion()
				discussion.DiscussableType = reflecth.GetTypeName(lecture)
				discussion.DiscussableId = lecture.Id
				discussion.DiscusserId = sliceh.Random(users...).Id
				discussion.DiscussionRepliedId = uuid.NullUUID{UUID: uuid.Nil, Valid: false}
				syncronizer.M().Lock()
				discussions = append(discussions, &discussion)
				syncronizer.M().Unlock()
			}
		}(lecture)
	}
	syncronizer.WG().Wait()
	for _, chunk := range sliceh.Chunk(discussions, 1000) {
		_, err = seeder.repository.Insert(c, chunk...)
	}

	return
}
