package models

import (
	"time"

	"github.com/google/uuid"
)

type PermissionRole struct {
	PermissionId uuid.UUID  `json:"permission_id" gorm:"type:CHAR(36);NOT NULL"`
	Permission   Permission `json:"permission" gorm:"foreignKey:PermissionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleId       uuid.UUID  `json:"role_id" gorm:"type:CHAR(36);NOT NULL"`
	Role         Role       `json:"role" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time  `json:"created_at"`
}

func (PermissionRole) TableName() string {
	return "permission_role"
}
