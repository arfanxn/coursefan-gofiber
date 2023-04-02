package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository instantiates a new PermissionRepository
func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

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
