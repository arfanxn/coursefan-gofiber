package repositories

import (
	"fmt"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

/*
 *	Repository instantiate method ⬇️
 */

// NewPermissionRepository instantiates a new PermissionRepository
func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

/*
 *	Repository utility methods ⬇️
 */

// GetModel returns the repository associated model
func (PermissionRepository) GetModel() models.Permission {
	return models.Permission{}
}

/*
 *	Repository query methods ⬇️
 */

// All returns all permissions in the database
func (repository *PermissionRepository) All(c *fiber.Ctx) (permissions []models.Permission, err error) {
	err = repository.db.Find(&permissions).Error
	return
}

// Find finds model by id
func (repository *PermissionRepository) Find(c *fiber.Ctx, id string) (permission models.Permission, err error) {
	err = repository.db.Where("id = ?", id).First(&permission).Error
	return
}

// Insert inserts models into the database
func (repository *PermissionRepository) Insert(c *fiber.Ctx, permissions ...*models.Permission) (int64, error) {
	for _, permission := range permissions {
		if permission.Id == uuid.Nil {
			permission.Id = uuid.New()
		}
	}
	result := repository.db.Create(permissions)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *PermissionRepository) UpdateById(c *fiber.Ctx, permission *models.Permission) (int64, error) {
	result := repository.db.Model(permission).Updates(permission)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *PermissionRepository) DeleteByIds(c *fiber.Ctx, permissions ...*models.Permission) (int64, error) {
	result := repository.db.Delete(permissions)
	return result.RowsAffected, result.Error
}

/*
 *	Repository permission name finder methods ⬇️
 */

// FindByNameAndCUR retrieves permission by name and by the CourseUserRole model from the given context.
func (repository *PermissionRepository) FindByNameAndCUR(
	c *fiber.Ctx,
	permissionName string,
) (models.Permission, error) {
	var (
		tx            = repository.db.Model(&models.Permission{})
		user          = ctxh.MustGetUser(c)
		permission    models.Permission
		courseId      string = c.Params("course_id")
		lecturePartId string = c.Params("lecture_part_id")
		lectureId     string = c.Params("lecture_id")
		reviewId      string = c.Params("review_id")
		discussionId  string = c.Params("discussion_id")
	)

	tx = tx.
		Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.PermissionRole{}.TableName(),
			models.PermissionRole{}.TableName(),
			"permission_id",
			models.Permission{}.TableName(),
			"id",
		)).
		Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.Role{}.TableName(),
			models.Role{}.TableName(),
			"id",
			models.PermissionRole{}.TableName(),
			"role_id",
		)).
		Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.CourseUserRole{}.TableName(),
			models.CourseUserRole{}.TableName(),
			"role_id",
			models.Role{}.TableName(),
			"id",
		)).
		Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.Course{}.TableName(),
			models.Course{}.TableName(),
			"id",
			models.CourseUserRole{}.TableName(),
			"course_id",
		))
	// If the lecture part id or lecture id is not empty, join the lecture table to check permissions
	if lecturePartId != "" || lectureId != "" {
		tx = tx.Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.LecturePart{}.TableName(),
			models.LecturePart{}.TableName(),
			"course_id",
			models.Course{}.TableName(),
			"id",
		))
	}
	// If the lecture id is not empty, join the lecture table to check permissions
	if lectureId != "" {
		tx = tx.Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.Lecture{}.TableName(),
			models.Lecture{}.TableName(),
			"lecture_part_id",
			models.LecturePart{}.TableName(),
			"id",
		))
	}
	// If the review id is not empty, join the review table to check permissions
	if reviewId != "" {
		tx = tx.Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.Review{}.TableName(),
			models.Review{}.TableName(),
			"reviewer_id",
			models.CourseUserRole{}.TableName(),
			"user_id",
		))
	}
	// If the discussion id is not empty, join the discussion table to check permissions
	if discussionId != "" {
		tx = tx.Joins(fmt.Sprintf("INNER JOIN %s ON %s.%s = %s.%s",
			models.Discussion{}.TableName(),
			models.Discussion{}.TableName(),
			"discusser_id",
			models.CourseUserRole{}.TableName(),
			"user_id",
		))
	}
	tx = tx.
		Where(models.Permission{}.TableName()+".name = ?", permissionName).
		Where(models.CourseUserRole{}.TableName()+".user_id = ?", user.Id.String())

	if courseId != "" {
		tx = tx.Where(models.Course{}.TableName()+".id = ?", courseId)
	}
	if lecturePartId != "" {
		tx = tx.Where(models.LecturePart{}.TableName()+".id = ?", lecturePartId)
	}
	if lectureId != "" {
		tx = tx.Where(models.Lecture{}.TableName()+".id = ?", lectureId)
	}
	if reviewId != "" {
		tx = tx.Where(models.Review{}.TableName()+".id = ?", reviewId)
	}
	if discussionId != "" {
		tx = tx.Where(models.Discussion{}.TableName()+".id = ?", discussionId)
	}

	tx = tx.First(&permission)
	err := tx.Error
	return permission, err
}
