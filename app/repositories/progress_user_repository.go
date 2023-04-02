package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgressUserRepository struct {
	db *gorm.DB
}

// NewProgressUserRepository instantiates a new ProgressUserRepository
func NewProgressUserRepository(db *gorm.DB) *ProgressUserRepository {
	return &ProgressUserRepository{db: db}
}

// All returns all progressUsers in the database
func (repository *ProgressUserRepository) All(c *fiber.Ctx) (progressUsers []models.ProgressUser, err error) {
	err = repository.db.Find(&progressUsers).Error
	return
}

// Find finds model by id
func (repository *ProgressUserRepository) Find(c *fiber.Ctx, id string) (progressUser models.ProgressUser, err error) {
	err = repository.db.Where("id = ?", id).First(&progressUser).Error
	return
}

// Insert inserts models into the database
func (repository *ProgressUserRepository) Insert(c *fiber.Ctx, progressUsers ...*models.ProgressUser) (int64, error) {
	for _, progressUser := range progressUsers {
		if progressUser.Id == uuid.Nil {
			progressUser.Id = uuid.New()
		}
	}
	result := repository.db.Create(progressUsers)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *ProgressUserRepository) UpdateById(c *fiber.Ctx, progressUser *models.ProgressUser) (int64, error) {
	result := repository.db.Model(progressUser).Updates(progressUser)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *ProgressUserRepository) DeleteByIds(c *fiber.Ctx, progressUsers ...*models.ProgressUser) (int64, error) {
	result := repository.db.Delete(progressUsers)
	return result.RowsAffected, result.Error
}
