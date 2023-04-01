package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type ProgressUser struct {
	Id         uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	ProgressId uuid.UUID `json:"progress_id" gorm:"type:CHAR(36);NOT NULL"`
	Progress   Progress  `json:"progress" gorm:"foreignKey:ProgressId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId     uuid.UUID `json:"user_id" gorm:"type:CHAR(36);NOT NULL"`
	User       User      `json:"user" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt  null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (ProgressUser) TableName() string {
	return "progress_user"
}
