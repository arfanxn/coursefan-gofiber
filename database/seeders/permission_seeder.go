package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type PermissionSeeder struct {
	repository *repositories.PermissionRepository
}

// NewPermissionSeeder instantiates a new PermissionSeeder
func NewPermissionSeeder(
	repository *repositories.PermissionRepository,
) *PermissionSeeder {
	return &PermissionSeeder{
		repository: repository,
	}
}

// Run runs the seeder
func (seeder *PermissionSeeder) Run(c *fiber.Ctx) (err error) {
	// Permission names
	permissionNames := enums.PermissionNames()
	// Seed
	var permissions []*models.Permission
	for _, permissionName := range permissionNames {
		permission := models.Permission{
			Name: permissionName,
		}
		permissions = append(permissions, &permission)

	}
	_, err = seeder.repository.Insert(c, permissions...)

	return
}
