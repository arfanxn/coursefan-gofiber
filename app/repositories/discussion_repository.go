package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscussionRepository struct {
	db *gorm.DB
}

// NewDiscussionRepository instantiates a new DiscussionRepository
func NewDiscussionRepository(db *gorm.DB) *DiscussionRepository {
	return &DiscussionRepository{db: db}
}

// All returns all discussions in the database
func (repository *DiscussionRepository) All(c *fiber.Ctx) (discussions []models.Discussion, err error) {
	err = repository.db.Find(&discussions).Error
	return
}

// Find finds model by id
func (repository *DiscussionRepository) Find(c *fiber.Ctx, id string) (discussion models.Discussion, err error) {
	err = repository.db.Where("id = ?", id).First(&discussion).Error
	return
}

// Insert inserts models into the database
func (repository *DiscussionRepository) Insert(c *fiber.Ctx, discussions ...*models.Discussion) (int64, error) {
	for _, discussion := range discussions {
		if discussion.Id == uuid.Nil {
			discussion.Id = uuid.New()
		}
	}
	result := repository.db.Create(discussions)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *DiscussionRepository) UpdateById(c *fiber.Ctx, discussion *models.Discussion) (int64, error) {
	result := repository.db.Model(discussion).Updates(discussion)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *DiscussionRepository) DeleteByIds(c *fiber.Ctx, discussions ...*models.Discussion) (int64, error) {
	result := repository.db.Delete(discussions)
	return result.RowsAffected, result.Error
}
