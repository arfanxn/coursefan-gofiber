package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PartRepository struct {
	db *gorm.DB
}

// NewPartRepository instantiates a new PartRepository
func NewPartRepository(db *gorm.DB) *PartRepository {
	return &PartRepository{db: db}
}

// All returns all parts in the database
func (repository *PartRepository) All(c *fiber.Ctx) (parts []models.Part, err error) {
	err = repository.db.Find(&parts).Error
	return
}

// Find finds model by id
func (repository *PartRepository) Find(c *fiber.Ctx, id string) (part models.Part, err error) {
	err = repository.db.Where("id = ?", id).First(&part).Error
	return
}

// Insert inserts models into the database
func (repository *PartRepository) Insert(c *fiber.Ctx, parts ...*models.Part) (int64, error) {
	for _, part := range parts {
		if part.Id == uuid.Nil {
			part.Id = uuid.New()
		}
	}
	result := repository.db.Create(parts)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *PartRepository) UpdateById(c *fiber.Ctx, part *models.Part) (int64, error) {
	result := repository.db.Model(part).Updates(part)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *PartRepository) DeleteByIds(c *fiber.Ctx, parts ...*models.Part) (int64, error) {
	result := repository.db.Delete(parts)
	return result.RowsAffected, result.Error
}
