package models

import "github.com/google/uuid"

type Permission struct {
	Id   uuid.UUID `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	Name string    `json:"name" gorm:"index;type:VARCHAR(50) NOT NULL"`
}

func (Permission) TableName() string {
	return "permissions"
}
