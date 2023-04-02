package seeders

import (
	"math/rand"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type LectureSeeder struct {
	repository            *repositories.LectureRepository
	lecturePartRepository *repositories.LecturePartRepository
}

// NewLectureSeeder instantiates a new LectureSeeder
func NewLectureSeeder(
	repository *repositories.LectureRepository,
	lecturePartRepository *repositories.LecturePartRepository,
) *LectureSeeder {
	return &LectureSeeder{
		repository:            repository,
		lecturePartRepository: lecturePartRepository,
	}
}

// Run runs the seeder
func (seeder *LectureSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	// Get all lecture parts
	lectureParts, err := seeder.lecturePartRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var lectures []*models.Lecture
	for _, lecturePart := range lectureParts {
		syncronizer.WG().Add(1)
		go func(lecturePart models.LecturePart) {
			defer syncronizer.WG().Done()
			for i := 0; i < rand.Intn(20); i++ {
				lecture := factories.FakeLecture()
				lecture.LecturePartId = lecturePart.Id
				lecture.Order = (i + 1)
				syncronizer.M().Lock()
				lectures = append(lectures, &lecture)
				syncronizer.M().Unlock()
			}
		}(lecturePart)
	}
	syncronizer.WG().Wait()
	_, err = seeder.repository.Insert(c, lectures...)

	return
}
