package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type ProgressSeeder struct {
	repository        *repositories.ProgressRepository
	lectureRepository *repositories.LectureRepository
}

// NewProgressSeeder instantiates a new ProgressSeeder
func NewProgressSeeder(
	repository *repositories.ProgressRepository,
	lectureRepository *repositories.LectureRepository,
) *ProgressSeeder {
	return &ProgressSeeder{
		repository:        repository,
		lectureRepository: lectureRepository,
	}
}

// Run runs the seeder
func (seeder *ProgressSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	// get all lectures
	lectures, err := seeder.lectureRepository.All(c)
	if err != nil {
		return
	}
	var progresses []*models.Progress
	for _, lecture := range lectures {
		progress := factories.FakeProgress()
		progress.ProgressableType = reflecth.GetTypeName(lecture)
		progress.ProgressableId = lecture.Id
		progresses = append(progresses, &progress)
	}

	_, err = seeder.repository.Insert(c, progresses...)

	return
}
