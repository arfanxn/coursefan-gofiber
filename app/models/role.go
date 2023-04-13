package models

import "github.com/google/uuid"

type Role struct {
	Id   uuid.UUID `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	Name string    `json:"name" gorm:"index;type:VARCHAR(25) NOT NULL"`
}

func (Role) TableName() string {
	return "roles"
}
