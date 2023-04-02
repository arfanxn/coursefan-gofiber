package seeders

import (
	"math/rand"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type LecturePartSeeder struct {
	repository       *repositories.LecturePartRepository
	courseRepository *repositories.CourseRepository
}

// NewLecturePartSeeder instantiates a new LecturePartSeeder
func NewLecturePartSeeder(
	repository *repositories.LecturePartRepository,
	courseRepository *repositories.CourseRepository,
) *LecturePartSeeder {
	return &LecturePartSeeder{
		repository:       repository,
		courseRepository: courseRepository,
	}
}

// Run runs the seeder
func (seeder *LecturePartSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	// Get all courses
	courses, err := seeder.courseRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var lectureParts []*models.LecturePart
	for _, course := range courses {
		syncronizer.WG().Add(1)
		go func(course models.Course) {
			defer syncronizer.WG().Done()
			for i := 0; i < rand.Intn(10); i++ {
				lecturePart := factories.FakeLecturePart()
				lecturePart.CourseId = course.Id
				lecturePart.Part = (i + 1)
				syncronizer.M().Lock()
				lectureParts = append(lectureParts, &lecturePart)
				syncronizer.M().Unlock()
			}
		}(course)
	}
	syncronizer.WG().Wait()
	_, err = seeder.repository.Insert(c, lectureParts...)

	return
}
