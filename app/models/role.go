package models

import "github.com/google/uuid"

type Role struct {
	Id   uuid.UUID `json:"id" gorm:"primaryKey"`
	Name string    `json:"name" gorm:"index;type:VARCHAR(25) NOT NULL"`
}
