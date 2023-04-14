package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
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
	courseWishlisterRole, err := seeder.roleRepository.FindByName(c, enums.RoleNameCourseWishlister)
	if err != nil {
		return
	}
	courseCarterRole, err := seeder.roleRepository.FindByName(c, enums.RoleNameCourseCarter)
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
			var curs []*models.CourseUserRole
			totalEachRelationKind := (len(users) / len(enums.RoleNames())) - 1

			shuffledUsers := sliceh.Shuffle(users)
			courseLecturerUser := shuffledUsers[0]
			courseParticipantUsersStartIdx := 1
			courseParticipantUsers := shuffledUsers[courseParticipantUsersStartIdx:(courseParticipantUsersStartIdx + totalEachRelationKind)]
			courseWishlisterUsersStartIdx := (len(courseParticipantUsers) + 1)
			courseWishlisterUsers := shuffledUsers[courseWishlisterUsersStartIdx:(courseWishlisterUsersStartIdx + totalEachRelationKind)]
			courseCarterUsersStartIdx := (len(courseParticipantUsers) + len(courseWishlisterUsers) + 1)
			courseCarterUsers := shuffledUsers[courseCarterUsersStartIdx:(courseCarterUsersStartIdx + totalEachRelationKind)]

			cur := factories.FakeCourseUserRole()
			cur.CourseId = course.Id
			cur.RoleId = courseLecturerRole.Id
			cur.UserId = courseLecturerUser.Id
			curs = append(curs, &cur)

			for _, user := range courseParticipantUsers {
				cur := factories.FakeCourseUserRole()
				cur.CourseId = course.Id
				cur.RoleId = courseParticipantRole.Id
				cur.UserId = user.Id
				curs = append(curs, &cur)
			}

			for _, user := range courseWishlisterUsers {
				cur := factories.FakeCourseUserRole()
				cur.CourseId = course.Id
				cur.RoleId = courseWishlisterRole.Id
				cur.UserId = user.Id
				curs = append(curs, &cur)
			}

			for _, user := range courseCarterUsers {
				cur := factories.FakeCourseUserRole()
				cur.CourseId = course.Id
				cur.RoleId = courseCarterRole.Id
				cur.UserId = user.Id
				curs = append(curs, &cur)
			}

			syncronizer.M().Lock()
			defer syncronizer.M().Unlock()
			courseUserRoleModels = append(courseUserRoleModels, curs...)
		}(course)
	}
	syncronizer.WG().Wait()
	if err = syncronizer.Err(); err != nil {
		return
	}
	for _, chunk := range sliceh.Chunk(courseUserRoleModels, 1000) {
		_, err = seeder.repository.Insert(c, chunk...)
	}

	return
}
