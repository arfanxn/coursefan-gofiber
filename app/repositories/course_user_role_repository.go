package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type CourseUserRoleRepository struct {
	db *gorm.DB
}

// NewCourseUserRoleRepository instantiates a new CourseUserRoleRepository
func NewCourseUserRoleRepository(db *gorm.DB) *CourseUserRoleRepository {
	return &CourseUserRoleRepository{db: db}
}

// All returns all courseUserRoles in the database
func (repository *CourseUserRoleRepository) All(c *fiber.Ctx) (courseUserRoles []models.CourseUserRole, err error) {
	err = repository.db.Find(&courseUserRoles).Error
	return
}

// Find finds model by id
func (repository *CourseUserRoleRepository) Find(c *fiber.Ctx, id string) (courseUserRole models.CourseUserRole, err error) {
	err = repository.db.Where("id = ?", id).First(&courseUserRole).Error
	return
}

// Insert inserts models into the database
func (repository *CourseUserRoleRepository) Insert(c *fiber.Ctx, courseUserRoles ...*models.CourseUserRole) (int64, error) {
	for _, courseUserRole := range courseUserRoles {
		if courseUserRole.Id == uuid.Nil {
			courseUserRole.Id = uuid.New()
		}
		courseUserRole.CreatedAt = time.Now()
	}
	result := repository.db.Create(courseUserRoles)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *CourseUserRoleRepository) UpdateById(c *fiber.Ctx, courseUserRole *models.CourseUserRole) (int64, error) {
	// refresh model updated at
	courseUserRole.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(courseUserRole).Updates(courseUserRole)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *CourseUserRoleRepository) DeleteByIds(c *fiber.Ctx, courseUserRoles ...*models.CourseUserRole) (int64, error) {
	result := repository.db.Delete(courseUserRoles)
	return result.RowsAffected, result.Error
}
