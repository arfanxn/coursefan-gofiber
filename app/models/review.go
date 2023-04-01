package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Review struct {
	Id             uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	ReviewableType string    `json:"reviewable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	ReviewableId   uuid.UUID `json:"reviewable_id" gorm:"type:VARCHAR(36) NOT NULL"`
	ReviewerId     uuid.UUID `json:"reviewer_id" gorm:"type:CHAR(36);NOT NULL"`
	Reviewer       User      `json:"reviewer" gorm:"foreignKey:ReviewerId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rate           int       `json:"collection_name"  gorm:"type:TINYINT UNSIGNED NOT NULL"`
	Title          string    `json:"title" gorm:"type:VARCHAR(50) NOT NULL"`
	Body           string    `json:"body" gorm:"type:VARCHAR(256) NOT NULL"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt      null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
