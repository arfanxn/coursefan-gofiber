package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

// NewReviewRepository instantiates a new ReviewRepository
func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

// All returns all reviews in the database
func (repository *ReviewRepository) All(c *fiber.Ctx) (reviews []models.Review, err error) {
	err = repository.db.Find(&reviews).Error
	return
}

// Find finds model by id
func (repository *ReviewRepository) Find(c *fiber.Ctx, id string) (review models.Review, err error) {
	err = repository.db.Where("id = ?", id).First(&review).Error
	return
}

// FindById finds model by id
func (repository *ReviewRepository) FindById(c *fiber.Ctx, id string) (review models.Review, err error) {
	err = repository.db.Where("id = ?", id).First(&review).Error
	return
}

// Insert inserts models into the database
func (repository *ReviewRepository) Insert(c *fiber.Ctx, reviews ...*models.Review) (int64, error) {
	for _, review := range reviews {
		if review.Id == uuid.Nil {
			review.Id = uuid.New()
		}
	}
	result := repository.db.Create(reviews)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *ReviewRepository) UpdateById(c *fiber.Ctx, review *models.Review) (int64, error) {
	result := repository.db.Model(review).Updates(review)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *ReviewRepository) DeleteByIds(c *fiber.Ctx, reviews ...*models.Review) (int64, error) {
	result := repository.db.Delete(reviews)
	return result.RowsAffected, result.Error
}
