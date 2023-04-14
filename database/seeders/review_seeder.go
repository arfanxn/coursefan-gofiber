package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type ReviewSeeder struct {
	repository     *repositories.ReviewRepository
	curRepository  *repositories.CourseUserRoleRepository
	roleRepository *repositories.RoleRepository
}

// NewReviewSeeder instantiates a new ReviewSeeder
func NewReviewSeeder(
	repository *repositories.ReviewRepository,
	curRepository *repositories.CourseUserRoleRepository,
	roleRepository *repositories.RoleRepository,
) *ReviewSeeder {
	return &ReviewSeeder{
		repository:     repository,
		curRepository:  curRepository,
		roleRepository: roleRepository,
	}
}

// Run runs the seeder
func (seeder *ReviewSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()

	courseUserRoleModels, err := seeder.curRepository.All(c)
	if err != nil {
		return
	}

	courseParticipantRole, err := seeder.roleRepository.FindByName(c, enums.RoleNameCourseParticipant)
	if err != nil {
		return
	}

	// Filter CourseUserRole Models only if CourseUserRole.Relation field is participant
	var participantCourseUserRoleModels []models.CourseUserRole
	for _, cur := range courseUserRoleModels {
		syncronizer.WG().Add(1)
		go func(cur models.CourseUserRole) {
			defer syncronizer.WG().Done()
			if cur.RoleId == courseParticipantRole.Id {
				syncronizer.M().Lock()
				participantCourseUserRoleModels = append(participantCourseUserRoleModels, cur)
				syncronizer.M().Unlock()
			}
		}(cur)
	}
	syncronizer.WG().Wait()

	// Seed
	var reviews []*models.Review
	for _, cur := range participantCourseUserRoleModels {
		syncronizer.WG().Add(1)
		go func(cur models.CourseUserRole) {
			defer syncronizer.WG().Done()
			review := factories.FakeReview()
			review.ReviewableType = reflecth.GetTypeName(models.Course{})
			review.ReviewableId = cur.CourseId
			review.ReviewerId = cur.UserId
			syncronizer.M().Lock()
			reviews = append(reviews, &review)
			syncronizer.M().Unlock()
		}(cur)
	}
	syncronizer.WG().Wait()
	_, err = seeder.repository.Insert(c, reviews...)

	return
}
