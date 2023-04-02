package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type RoleSeeder struct {
	repository *repositories.RoleRepository
}

// NewRoleSeeder instantiates a new RoleSeeder
func NewRoleSeeder(
	repository *repositories.RoleRepository,
) *RoleSeeder {
	return &RoleSeeder{
		repository: repository,
	}
}

// Run runs the seeder
func (seeder *RoleSeeder) Run(c *fiber.Ctx) (err error) {
	// Role names
	roleNames := enums.RoleNames()
	// Seed
	var roles []*models.Role
	for _, roleName := range roleNames {
		role := models.Role{
			Name: roleName,
		}
		roles = append(roles, &role)

	}
	_, err = seeder.repository.Insert(c, roles...)

	return
}
