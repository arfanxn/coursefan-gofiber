package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Part struct {
	Id           uuid.UUID   `json:"id" gorm:"primaryKey;type:char(36)"`
	PartableType string      `json:"partable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	PartableId   uuid.UUID   `json:"partable_id" gorm:"type:VARCHAR(36) NOT NULL"`
	Part         int         `json:"part"  gorm:"type:INTEGER UNSIGNED NOT NULL"`
	Name         string      `json:"title" gorm:"type:VARCHAR(50) NOT NULL"`
	Description  null.String `json:"description" gorm:"type:VARCHAR(256)"`
	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt    null.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}
