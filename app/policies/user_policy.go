package policies

import "github.com/arfanxn/coursefan-gofiber/app/repositories"

type UserPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewUserPolicy(permissionRepository *repositories.PermissionRepository) *UserPolicy {
	return &UserPolicy{permissionRepository: permissionRepository}
}
