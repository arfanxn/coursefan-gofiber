package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type PermissionRoleRepository struct {
	db *gorm.DB
}

// NewPermissionRoleRepository instantiates a new PermissionRoleRepository
func NewPermissionRoleRepository(db *gorm.DB) *PermissionRoleRepository {
	return &PermissionRoleRepository{db: db}
}

// All returns all permissionRoles in the database
func (repository *PermissionRoleRepository) All(c *fiber.Ctx) (permissionRoles []models.PermissionRole, err error) {
	err = repository.db.Find(&permissionRoles).Error
	return
}

// Find finds model by id
func (repository *PermissionRoleRepository) Find(c *fiber.Ctx, id string) (permissionRole models.PermissionRole, err error) {
	err = repository.db.Where("id = ?", id).First(&permissionRole).Error
	return
}

// Insert inserts models into the database
func (repository *PermissionRoleRepository) Insert(c *fiber.Ctx, permissionRoles ...*models.PermissionRole) (int64, error) {
	for _, permissionRole := range permissionRoles {
		if permissionRole.Id == uuid.Nil {
			permissionRole.Id = uuid.New()
		}
		permissionRole.CreatedAt = time.Now()
	}
	result := repository.db.Create(permissionRoles)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *PermissionRoleRepository) UpdateById(c *fiber.Ctx, permissionRole *models.PermissionRole) (int64, error) {
	// refresh model updated at
	permissionRole.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(permissionRole).Updates(permissionRole)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *PermissionRoleRepository) DeleteByIds(c *fiber.Ctx, permissionRoles ...*models.PermissionRole) (int64, error) {
	result := repository.db.Delete(permissionRoles)
	return result.RowsAffected, result.Error
}
