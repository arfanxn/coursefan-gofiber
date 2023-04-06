package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type LecturePartRepository struct {
	db *gorm.DB
}

// NewLecturePartRepository instantiates a new LecturePartRepository
func NewLecturePartRepository(db *gorm.DB) *LecturePartRepository {
	return &LecturePartRepository{db: db}
}

// All returns all lectureParts in the database
func (repository *LecturePartRepository) All(c *fiber.Ctx) (lectureParts []models.LecturePart, err error) {
	err = repository.db.Find(&lectureParts).Error
	return
}

// Find finds model by id
func (repository *LecturePartRepository) Find(c *fiber.Ctx, id string) (lecturePart models.LecturePart, err error) {
	err = repository.db.Where("id = ?", id).First(&lecturePart).Error
	return
}

// Insert inserts models into the database
func (repository *LecturePartRepository) Insert(c *fiber.Ctx, lectureParts ...*models.LecturePart) (int64, error) {
	for _, lecturePart := range lectureParts {
		if lecturePart.Id == uuid.Nil {
			lecturePart.Id = uuid.New()
		}
		lecturePart.CreatedAt = time.Now()
	}
	result := repository.db.Create(lectureParts)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *LecturePartRepository) UpdateById(c *fiber.Ctx, lecturePart *models.LecturePart) (int64, error) {
	// refresh model updated at
	lecturePart.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(lecturePart).Updates(lecturePart)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *LecturePartRepository) DeleteByIds(c *fiber.Ctx, lectureParts ...*models.LecturePart) (int64, error) {
	result := repository.db.Delete(lectureParts)
	return result.RowsAffected, result.Error
}