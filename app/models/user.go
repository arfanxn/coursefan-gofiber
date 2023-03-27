package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID    `json:"id" gorm:"primaryKey"`
	Name      string       `json:"name" gorm:"index"`
	Email     string       `json:"email" gorm:"index"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}
