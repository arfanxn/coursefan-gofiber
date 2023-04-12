package repositories

import (
	"fmt"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
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
func (repository *LecturePartRepository) All(c *fiber.Ctx, queries ...requests.Query) (
	lectureParts []models.LecturePart, err error) {
	db := repository.db
	if len(queries) != 0 {
		db = gormh.BuildFromRequestQuery(repository.db, models.LecturePart{}, queries[0])
	}
	err = db.Find(&lectureParts).Error
	return
}

// AllByCourse returns all lectureParts by course
func (repository *LecturePartRepository) AllByCourse(c *fiber.Ctx, query requests.Query) (
	lectureParts []models.LecturePart, err error) {
	courseIdFilter := query.GetFilter("lecture_parts.course_id", enums.QueryFilterOperatorEquals)
	err = gormh.BuildFromRequestQuery(repository.db, models.LecturePart{}, query).
		Joins(
			fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
				models.Course{}.TableName(),
				models.Course{}.TableName(),
				"id",
				models.LecturePart{}.TableName(),
				"course_id",
			)).
		Joins(
			fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
				models.CourseUserRole{}.TableName(),
				models.CourseUserRole{}.TableName(),
				"course_id",
				models.Course{}.TableName(),
				"id",
			)).
		Where(models.Course{}.TableName()+".id = ?", courseIdFilter.Values[0]).
		Distinct().Find(&lectureParts).Error

	return
}

// FindById finds model by id
func (repository *LecturePartRepository) FindById(c *fiber.Ctx, id string) (lecturePart models.LecturePart, err error) {
	err = repository.db.Where("id = ?", id).First(&lecturePart).Error
	return
}

// FindByModel finds model by model
func (repository *LecturePartRepository) FindByModel(c *fiber.Ctx, model models.LecturePart) (lecturePart models.LecturePart, err error) {
	err = repository.db.First(&lecturePart, model).Error
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
