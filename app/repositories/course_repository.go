package repositories

import (
	"fmt"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

// NewCourseRepository instantiates a new CourseRepository
func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

// All returns all courses in the database, queries argument is optional
func (repository *CourseRepository) All(c *fiber.Ctx, queries ...requests.Query) (
	courses []models.Course, err error) {
	db := repository.db
	if len(queries) != 0 {
		query := queries[0]
		db = gormh.BuildFromRequestQuery(repository.db, models.Course{}, query)

		if query.HasScope(
			enums.CourseQueryScopeLectured,
			enums.CourseQueryScopeParticipated,
			enums.CourseQueryScopeCart,
			enums.CourseQueryScopeWishlist,
		) {
			db = db.Joins(
				fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
					models.CourseUserRole{}.TableName(),
					models.CourseUserRole{}.TableName(),
					"course_id",
					models.Course{}.TableName(),
					"id",
				)).Joins(
				fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
					models.Role{}.TableName(),
					models.Role{}.TableName(),
					"id",
					models.CourseUserRole{}.TableName(),
					"role_id",
				)).
				Where(models.CourseUserRole{}.TableName()+".user_id", ctxh.MustGetUser(c).Id)

			if query.HasScope(enums.CourseQueryScopeLectured) {
				db = db.
					Where(models.Role{}.TableName()+".name", enums.RoleNameCourseLecturer)
			} else if query.HasScope(enums.CourseQueryScopeParticipated) {
				db = db.
					Where(models.Role{}.TableName()+".name", enums.RoleNameCourseParticipant)
			} else if query.HasScope(enums.CourseQueryScopeCart) {
				db = db.
					Where(models.Role{}.TableName()+".name", enums.RoleNameCourseCarter)
			} else if query.HasScope(enums.CourseQueryScopeWishlist) {
				db = db.
					Where(models.Role{}.TableName()+".name", enums.RoleNameCourseWishlister)
			}
		}
	}
	err = db.Distinct().Find(&courses).Error
	return
}

// FindById finds model by id
func (repository *CourseRepository) FindById(c *fiber.Ctx, id string) (course models.Course, err error) {
	err = repository.db.Where("id = ?", id).First(&course).Error
	return
}

// FindBySlug finds model by slug
func (repository *CourseRepository) FindBySlug(c *fiber.Ctx, id string) (course models.Course, err error) {
	err = repository.db.Where("slug = ?", id).First(&course).Error
	return
}

// Insert inserts models into the database
func (repository *CourseRepository) Insert(c *fiber.Ctx, courses ...*models.Course) (int64, error) {
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
func (repository *CourseRepository) UpdateById(c *fiber.Ctx, course *models.Course) (int64, error) {
	// refresh model updated at
	course.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(course).Updates(course)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *CourseRepository) DeleteByIds(c *fiber.Ctx, courses ...*models.Course) (int64, error) {
	result := repository.db.Delete(courses)
	return result.RowsAffected, result.Error
}
