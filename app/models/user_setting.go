package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserSetting struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:CHAR(36);NOT NULL"`
	User      *User     `json:"user" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Key       string    `json:"key" gorm:"type:VARCHAR(50) NOT NULL"`
	Value     string    `json:"value" gorm:"type:VARCHAR(50) NOT NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (UserSetting) TableName() string {
	return "user_settings"
}
