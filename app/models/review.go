package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Review struct {
	Id             uuid.UUID   `json:"id" gorm:"primaryKey;type:char(36)"`
	ReviewableType string      `json:"reviewable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	ReviewableId   uuid.UUID   `json:"reviewable_id" gorm:"type:VARCHAR(36) NOT NULL"`
	ReviewerId     uuid.UUID   `json:"reviewer_id" gorm:"type:CHAR(36);NOT NULL"`
	Reviewer       User        `json:"reviewer" gorm:"foreignKey:ReviewerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rate           int         `json:"rate"  gorm:"type:TINYINT UNSIGNED NOT NULL"`
	Title          null.String `json:"title" gorm:"type:VARCHAR(50)"`
	Body           null.String `json:"body" gorm:"type:TEXT"`
	CreatedAt      time.Time   `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt      null.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Review) TableName() string {
	return "reviews"
}
