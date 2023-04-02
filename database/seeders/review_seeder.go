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
	repository    *repositories.ReviewRepository
	cusRepository *repositories.CourseUserRoleRepository
}

// NewReviewSeeder instantiates a new ReviewSeeder
func NewReviewSeeder(
	repository *repositories.ReviewRepository,
	cusRepository *repositories.CourseUserRoleRepository,
) *ReviewSeeder {
	return &ReviewSeeder{
		repository:    repository,
		cusRepository: cusRepository,
	}
}

// Run runs the seeder
func (seeder *ReviewSeeder) Run(c *fiber.Ctx) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()

	courseUserRoleModels, err := seeder.cusRepository.All(c)
	if err != nil {
		return
	}
	// Filter CourseUserRole Models only if CourseUserRole.Relation field is participant
	var participantCourseUserRoleModels []models.CourseUserRole
	for _, cus := range courseUserRoleModels {
		syncronizer.WG().Add(1)
		go func(cus models.CourseUserRole) {
			defer syncronizer.WG().Done()
			if cus.Relation == enums.CourseUserRoleRelationParticipant {
				syncronizer.M().Lock()
				participantCourseUserRoleModels = append(participantCourseUserRoleModels, cus)
				syncronizer.M().Unlock()
			}
		}(cus)
	}
	syncronizer.WG().Wait()

	// Seed
	var reviews []*models.Review
	for _, cus := range participantCourseUserRoleModels {
		syncronizer.WG().Add(1)
		go func(cus models.CourseUserRole) {
			defer syncronizer.WG().Done()
			review := factories.FakeReview()
			review.ReviewableType = reflecth.GetTypeName(models.Course{})
			review.ReviewableId = cus.CourseId
			review.ReviewerId = cus.UserId
			syncronizer.M().Lock()
			reviews = append(reviews, &review)
			syncronizer.M().Unlock()
		}(cus)
	}
	syncronizer.WG().Wait()
	_, err = seeder.repository.Insert(c, reviews...)

	return
}
