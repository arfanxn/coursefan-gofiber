package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
)

type UserSettingPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewUserSettingPolicy(permissionRepository *repositories.PermissionRepository) *UserSettingPolicy {
	return &UserSettingPolicy{permissionRepository: permissionRepository}
}
