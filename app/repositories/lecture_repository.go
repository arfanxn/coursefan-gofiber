package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type LectureRepository struct {
	db *gorm.DB
}

// NewLectureRepository instantiates a new LectureRepository
func NewLectureRepository(db *gorm.DB) *LectureRepository {
	return &LectureRepository{db: db}
}

// All returns all lectures in the database
func (repository *LectureRepository) All(c *fiber.Ctx) (lectures []models.Lecture, err error) {
	err = repository.db.Find(&lectures).Error
	return
}

// Find finds model by id
func (repository *LectureRepository) Find(c *fiber.Ctx, id string) (lecture models.Lecture, err error) {
	err = repository.db.Where("id = ?", id).First(&lecture).Error
	return
}

// Insert inserts models into the database
func (repository *LectureRepository) Insert(c *fiber.Ctx, lectures ...*models.Lecture) (int64, error) {
	for _, lecture := range lectures {
		if lecture.Id == uuid.Nil {
			lecture.Id = uuid.New()
		}
		lecture.CreatedAt = time.Now()
	}
	result := repository.db.Create(lectures)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *LectureRepository) UpdateById(c *fiber.Ctx, lecture *models.Lecture) (int64, error) {
	// refresh model updated at
	lecture.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(lecture).Updates(lecture)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *LectureRepository) DeleteByIds(c *fiber.Ctx, lectures ...*models.Lecture) (int64, error) {
	result := repository.db.Delete(lectures)
	return result.RowsAffected, result.Error
}
