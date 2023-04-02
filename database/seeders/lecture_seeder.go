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
	repository       *repositories.LectureRepository
	courseRepository *repositories.CourseRepository
}

// NewLectureSeeder instantiates a new LectureSeeder
func NewLectureSeeder(
	repository *repositories.LectureRepository,
	courseRepository *repositories.CourseRepository,
) *LectureSeeder {
	return &LectureSeeder{
		repository:       repository,
		courseRepository: courseRepository,
	}
}

// Run runs the seeder
func (seeder *LectureSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	// Get all courses
	courses, err := seeder.courseRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var lectures []*models.Lecture
	for _, course := range courses {
		syncronizer.WG().Add(1)
		go func(course models.Course) {
			defer syncronizer.WG().Done()
			for i := 0; i < rand.Intn(50); i++ {
				lecture := factories.FakeLecture()
				lecture.CourseId = course.Id
				lecture.Order = (i + 1)
				syncronizer.M().Lock()
				lectures = append(lectures, &lecture)
				syncronizer.M().Unlock()
			}
		}(course)
	}
	syncronizer.WG().Wait()
	_, err = seeder.repository.Insert(c, lectures...)

	return
}
