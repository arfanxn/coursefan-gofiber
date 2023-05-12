package policies

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
)

type WalletPolicy struct {
	permissionRepository *repositories.PermissionRepository
}

func NewWalletPolicy(permissionRepository *repositories.PermissionRepository) *WalletPolicy {
	return &WalletPolicy{permissionRepository: permissionRepository}
}
