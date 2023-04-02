package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type PermissionRoleSeeder struct {
	repository           *repositories.PermissionRoleRepository
	permissionRepository *repositories.PermissionRepository
	roleRepository       *repositories.RoleRepository
}

// NewPermissionRoleSeeder instantiates a new PermissionRoleSeeder
func NewPermissionRoleSeeder(
	repository *repositories.PermissionRoleRepository,
	permissionRepository *repositories.PermissionRepository,
	roleRepository *repositories.RoleRepository,
) *PermissionRoleSeeder {
	return &PermissionRoleSeeder{
		repository:           repository,
		permissionRepository: permissionRepository,
		roleRepository:       roleRepository,
	}
}

// Run runs the seeder
func (seeder *PermissionRoleSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all permissions
	permissions, err := seeder.permissionRepository.All(c)
	if err != nil {
		return
	}
	courseLecturerPermissions := permissions
	courseParticipantPermissions := sliceh.Filter(permissions, func(permission models.Permission) bool {
		return sliceh.Contains([]string{
			enums.PermissionNameCourseView,

			enums.PermissionNameCourseReviewView,
			enums.PermissionNameCourseReviewCreate,
			enums.PermissionNameCourseReviewEdit,
			enums.PermissionNameCourseReviewDelete,

			enums.PermissionNameLectureView,

			enums.PermissionNameLecturePartView,
		}, permission.Name)
	})
	// Get all roles
	roles, err := seeder.roleRepository.All(c)
	if err != nil {
		return
	}
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	// Seed
	var permissionRoles []*models.PermissionRole
	for _, role := range roles {
		syncronizer.WG().Add(1)
		go func(role models.Role) {
			defer syncronizer.WG().Done()
			var permissionRole_permissions []models.Permission
			var permissionRole models.PermissionRole
			permissionRole.RoleId = role.Id
			switch role.Name {
			case enums.RoleNameCourseLecturer:
				permissionRole_permissions = courseLecturerPermissions
				break
			case enums.RoleNameCourseParticipant:
				permissionRole_permissions = courseParticipantPermissions
				break
			}
			for _, permission := range permissionRole_permissions {
				permissionRole.PermissionId = permission.Id
				syncronizer.M().Lock()
				permissionRoles = append(permissionRoles, &permissionRole)
				syncronizer.M().Unlock()
			}
		}(role)
	}
	syncronizer.WG().Wait()
	if err = syncronizer.Err(); err != nil {
		return
	}
	_, err = seeder.repository.Insert(c, permissionRoles...)

	return
}
