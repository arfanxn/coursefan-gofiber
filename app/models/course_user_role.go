package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type CourseUserRole struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:CHAR(36);NOT NULL"`
	CourseId  uuid.UUID `json:"permission_id" gorm:"type:CHAR(36);NOT NULL"`
	Course    *Course    `json:"permission" gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:CHAR(36);NOT NULL"`
	User      *User     `json:"user" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleId    uuid.UUID `json:"role_id" gorm:"type:CHAR(36);NOT NULL"`
	Role      *Role     `json:"role" gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (CourseUserRole) TableName() string {
	return "course_user_role"
}
