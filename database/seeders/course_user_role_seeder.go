package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CourseUserRoleSeeder struct {
	repository       *repositories.CourseUserRoleRepository
	courseRepository *repositories.CourseRepository
	userRepository   *repositories.UserRepository
	roleRepository   *repositories.RoleRepository
}

// NewCourseUserRoleSeeder instantiates a new CourseUserRoleSeeder
func NewCourseUserRoleSeeder(
	repository *repositories.CourseUserRoleRepository,
	courseRepository *repositories.CourseRepository,
	userRepository *repositories.UserRepository,
	roleRepository *repositories.RoleRepository,
) *CourseUserRoleSeeder {
	return &CourseUserRoleSeeder{
		repository:       repository,
		courseRepository: courseRepository,
		userRepository:   userRepository,
		roleRepository:   roleRepository,
	}
}

// Run runs the seeder
func (seeder *CourseUserRoleSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all courses
	courses, err := seeder.courseRepository.All(c)
	if err != nil {
		return
	}
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Get roles
	courseLecturerRole, err := seeder.roleRepository.FindByName(c, enums.RoleNameCourseLecturer)
	if err != nil {
		return
	}
	courseParticipantRole, err := seeder.roleRepository.FindByName(c, enums.RoleNameCourseParticipant)
	if err != nil {
		return
	}
	// Seed
	syncronizer := synch.NewSyncronizer()
	var courseUserRoleModels []*models.CourseUserRole
	for _, course := range courses {
		syncronizer.WG().Add(1)
		go func(course models.Course) {
			defer syncronizer.WG().Done()
			var cuss []*models.CourseUserRole
			totalEachRelationKind := (len(users) / len(enums.CourseUserRoleRelations())) - 1

			shuffledUsers := sliceh.Shuffle(users)
			courseLecturerUser := shuffledUsers[0]
			courseParticipantUsers := shuffledUsers[1:totalEachRelationKind]
			courseWishlisterUsers := shuffledUsers[(len(courseParticipantUsers) + 1):totalEachRelationKind]
			courseCarterUsers := shuffledUsers[(len(courseParticipantUsers) + len(courseWishlisterUsers) + 1):totalEachRelationKind]

			cus := factories.FakeCourseUserRole()
			cus.CourseId = course.Id
			cus.RoleId = uuid.NullUUID{UUID: courseLecturerRole.Id, Valid: true}
			cus.UserId = courseLecturerUser.Id
			cuss = append(cuss, &cus)

			for _, user := range courseParticipantUsers {
				cus := factories.FakeCourseUserRole()
				cus.CourseId = course.Id
				cus.RoleId = uuid.NullUUID{UUID: courseParticipantRole.Id, Valid: true}
				cus.UserId = user.Id
				cuss = append(cuss, &cus)
			}

			for _, user := range courseWishlisterUsers {
				cus := factories.FakeCourseUserRole()
				cus.CourseId = course.Id
				cus.RoleId = uuid.NullUUID{UUID: uuid.UUID{}, Valid: false}
				cus.UserId = user.Id
				cuss = append(cuss, &cus)
			}

			for _, user := range courseCarterUsers {
				cus := factories.FakeCourseUserRole()
				cus.CourseId = course.Id
				cus.RoleId = uuid.NullUUID{UUID: uuid.UUID{}, Valid: false}
				cus.UserId = user.Id
				cuss = append(cuss, &cus)
			}

			syncronizer.M().Lock()
			defer syncronizer.M().Unlock()
			courseUserRoleModels = append(courseUserRoleModels, cuss...)
		}(course)
	}
	syncronizer.WG().Wait()
	if err = syncronizer.Err(); err != nil {
		return
	}
	_, err = seeder.repository.Insert(c, courseUserRoleModels...)

	return
}
