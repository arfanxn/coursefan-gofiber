package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	Name      string         `json:"name" gorm:"index;type:VARCHAR(50) NOT NULL"`
	Email     string         `json:"email" gorm:"uniqueIndex;type:VARCHAR(50) NOT NULL"`
	Password  string         `json:"password" gorm:"type:VARCHAR(256) NOT NULL"`
	Biography sql.NullString `json:"biography" gorm:"type:VARCHAR(256)"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt sql.NullTime   `json:"updated_at" gorm:"autoUpdateTime"`

	// Avatar relationship
	Avatar Media `json:"avatar" gorm:"-"`
}
