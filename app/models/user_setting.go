package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserSetting struct {
	Id        uuid.UUID    `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	UserId    uuid.UUID    `json:"user_id" gorm:"type:CHAR(36);NOT NULL"`
	User      User         `json:"user" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Key       string       `json:"key" gorm:"type:VARCHAR(50) NOT NULL"`
	Value     string       `json:"value" gorm:"type:VARCHAR(50) NOT NULL"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}
