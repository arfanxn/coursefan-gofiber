package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgressRepository struct {
	db *gorm.DB
}

// NewProgressRepository instantiates a new ProgressRepository
func NewProgressRepository(db *gorm.DB) *ProgressRepository {
	return &ProgressRepository{db: db}
}

// All returns all progresses in the database
func (repository *ProgressRepository) All(c *fiber.Ctx) (progresses []models.Progress, err error) {
	err = repository.db.Find(&progresses).Error
	return
}

// Find finds model by id
func (repository *ProgressRepository) Find(c *fiber.Ctx, id string) (progress models.Progress, err error) {
	err = repository.db.Where("id = ?", id).First(&progress).Error
	return
}

// Insert inserts models into the database
func (repository *ProgressRepository) Insert(c *fiber.Ctx, progresses ...*models.Progress) (int64, error) {
	for _, progress := range progresses {
		if progress.Id == uuid.Nil {
			progress.Id = uuid.New()
		}
	}
	result := repository.db.Create(progresses)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *ProgressRepository) UpdateById(c *fiber.Ctx, progress *models.Progress) (int64, error) {
	result := repository.db.Model(progress).Updates(progress)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *ProgressRepository) DeleteByIds(c *fiber.Ctx, progresses ...*models.Progress) (int64, error) {
	result := repository.db.Delete(progresses)
	return result.RowsAffected, result.Error
}
