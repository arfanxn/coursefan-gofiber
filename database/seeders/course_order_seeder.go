package seeders

import (
	"math/rand"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type CourseOrderSeeder struct {
	repository       *repositories.CourseOrderRepository
	courseRepository *repositories.CourseRepository
	userRepository   *repositories.UserRepository
}

// NewCourseOrderSeeder instantiates a new CourseOrderSeeder
func NewCourseOrderSeeder(
	repository *repositories.CourseOrderRepository,
	courseRepository *repositories.CourseRepository,
	userRepository *repositories.UserRepository,
) *CourseOrderSeeder {
	return &CourseOrderSeeder{
		repository:       repository,
		courseRepository: courseRepository,
		userRepository:   userRepository,
	}
}

// Run runs the seeder
func (seeder *CourseOrderSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Get all courses
	courses, err := seeder.courseRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var courseOrders []*models.CourseOrder
	for _, user := range users {
		course := courses[rand.Intn(len(courses))-1]
		courseOrder := factories.FakeCourseOrder()
		courseOrder.CourseId = course.Id
		courseOrder.UserId = user.Id
		courseOrders = append(courseOrders, &courseOrder)
	}
	_, err = seeder.repository.Insert(c, courseOrders...)

	return
}
