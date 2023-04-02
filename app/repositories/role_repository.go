package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository instantiates a new RoleRepository
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// All returns all roles in the database
func (repository *RoleRepository) All(c *fiber.Ctx) (roles []models.Role, err error) {
	err = repository.db.Find(&roles).Error
	return
}

// Find finds model by id
func (repository *RoleRepository) Find(c *fiber.Ctx, id string) (role models.Role, err error) {
	err = repository.db.Where("id = ?", id).First(&role).Error
	return
}

// Insert inserts models into the database
func (repository *RoleRepository) Insert(c *fiber.Ctx, roles ...*models.Role) (int64, error) {
	for _, role := range roles {
		if role.Id == uuid.Nil {
			role.Id = uuid.New()
		}
	}
	result := repository.db.Create(roles)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *RoleRepository) UpdateById(c *fiber.Ctx, role *models.Role) (int64, error) {
	result := repository.db.Model(role).Updates(role)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *RoleRepository) DeleteByIds(c *fiber.Ctx, roles ...*models.Role) (int64, error) {
	result := repository.db.Delete(roles)
	return result.RowsAffected, result.Error
}
