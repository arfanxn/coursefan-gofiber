package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type UserProfileRepository struct {
	db *gorm.DB
}

// NewUserProfileRepository instantiates a new UserProfileRepository
func NewUserProfileRepository(db *gorm.DB) *UserProfileRepository {
	return &UserProfileRepository{db: db}
}

// All returns all user profiles in the database
func (repository *UserProfileRepository) All(c *fiber.Ctx) (userProfiles []models.UserProfile, err error) {
	err = repository.db.Find(&userProfiles).Error
	return
}

// Find finds model by id
func (repository *UserProfileRepository) Find(c *fiber.Ctx, id string) (user models.User, err error) {
	err = repository.db.Where("id = ?", id).First(&user).Error
	return
}

// Insert inserts models into the database
func (repository *UserProfileRepository) Insert(c *fiber.Ctx, userProfiles ...*models.UserProfile) (int64, error) {
	for _, userProfile := range userProfiles {
		if userProfile.Id == uuid.Nil {
			userProfile.Id = uuid.New()
		}
		userProfile.CreatedAt = time.Now()
	}
	result := repository.db.Create(userProfiles)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *UserProfileRepository) UpdateById(c *fiber.Ctx, userProfile *models.UserProfile) (int64, error) {
	// refresh model updated at
	userProfile.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(userProfile).Updates(userProfile)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *UserProfileRepository) DeleteByIds(c *fiber.Ctx, userProfiles ...*models.UserProfile) (int64, error) {
	result := repository.db.Delete(userProfiles)
	return result.RowsAffected, result.Error
}
