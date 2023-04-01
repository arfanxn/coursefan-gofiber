package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Progress struct {
	Id               uuid.UUID `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	ProgressableType string    `json:"progressable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	ProgressableId   uuid.UUID `json:"progressable_id" gorm:"type:VARCHAR(36) NOT NULL"`
	Percentage       int       `json:"percetage" gorm:"type:TINYINT UNSIGNED NOT NULL;default:0;comment:this column represents the progress percentage"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt        null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
