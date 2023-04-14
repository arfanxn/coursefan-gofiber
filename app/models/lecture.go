package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Lecture struct {
	Id            uuid.UUID    `json:"id" gorm:"primaryKey;type:char(36)"`
	LecturePartId uuid.UUID    `json:"course_id" gorm:"type:CHAR(36);NOT NULL"`
	LecturePart   *LecturePart `json:"course" gorm:"foreignKey:LecturePartId;constraint:OnUpdate:CASCADE;"`
	Name          string       `json:"name" gorm:"type:VARCHAR(50);NOT NULL"`
	Order         int          `json:"order" gorm:"type:INTEGER UNSIGNED NOT NULL;default:0"`
	CreatedAt     time.Time    `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt     null.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Lecture) TableName() string {
	return "lectures"
}
