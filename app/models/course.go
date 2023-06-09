package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Course struct {
	Id           uuid.UUID     `json:"id" gorm:"primaryKey;type:char(36)"`
	Name         string        `json:"name" gorm:"type:VARCHAR(50);NOT NULL"`
	Slug         string        `json:"slug" gorm:"uniqueIndex;type:VARCHAR(50);NOT NULL"`
	Description  string        `json:"description" gorm:"type:LONGTEXT;NOT NULL"`
	Price        float64       `json:"price" gorm:"type:DECIMAL(10,2);NOT NULL"`
	CreatedAt    time.Time     `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt    null.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	LectureParts []LecturePart `json:"lecture_parts" gorm:"foreignKey:CourseId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Course) TableName() string {
	return "courses"
}
