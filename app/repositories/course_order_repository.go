package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type CourseOrderRepository struct {
	db *gorm.DB
}

// NewCourseOrderRepository instantiates a new CourseOrderRepository
func NewCourseOrderRepository(db *gorm.DB) *CourseOrderRepository {
	return &CourseOrderRepository{db: db}
}

// All returns all reviews in the database
func (repository *CourseOrderRepository) All(c *fiber.Ctx, queries ...requests.Query) (reviews []models.CourseOrder, err error) {
	tx := repository.db
	if query := sliceh.FirstOrNil(queries); query != nil {
		tx = gormh.BuildFromRequestQuery(repository.db, models.CourseOrder{}, *query)
	}
	err = tx.Find(&reviews).Error
	return
}

// FindById finds model by id
func (repository *CourseOrderRepository) FindById(c *fiber.Ctx, id string) (course models.CourseOrder, err error) {
	err = repository.db.Where("id = ?", id).First(&course).Error
	return
}

// FindBySlug finds model by slug
func (repository *CourseOrderRepository) FindBySlug(c *fiber.Ctx, id string) (course models.CourseOrder, err error) {
	err = repository.db.Where("slug = ?", id).First(&course).Error
	return
}

// Insert inserts models into the database
func (repository *CourseOrderRepository) Insert(c *fiber.Ctx, courses ...*models.CourseOrder) (int64, error) {
	for _, course := range courses {
		if course.Id == uuid.Nil {
			course.Id = uuid.New()
		}
		course.CreatedAt = time.Now()
	}
	result := repository.db.Create(courses)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *CourseOrderRepository) UpdateById(c *fiber.Ctx, course *models.CourseOrder) (int64, error) {
	// refresh model updated at
	course.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(course).Updates(course)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *CourseOrderRepository) DeleteByIds(c *fiber.Ctx, courses ...*models.CourseOrder) (int64, error) {
	result := repository.db.Delete(courses)
	return result.RowsAffected, result.Error
}
