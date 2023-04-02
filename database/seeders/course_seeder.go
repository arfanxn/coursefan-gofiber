package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type CourseSeeder struct {
	repository     *repositories.CourseRepository
	userRepository *repositories.UserRepository
}

// NewCourseSeeder instantiates a new CourseSeeder
func NewCourseSeeder(
	repository *repositories.CourseRepository,
	userRepository *repositories.UserRepository,
) *CourseSeeder {
	return &CourseSeeder{
		repository:     repository,
		userRepository: userRepository,
	}
}

// Run runs the seeder
func (seeder *CourseSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var courses []*models.Course
	for range users {
		for i := 0; i < 2; i++ {
			course := factories.FakeCourse()
			courses = append(courses, &course)
		}
	}
	_, err = seeder.repository.Insert(c, courses...)

	return
}
