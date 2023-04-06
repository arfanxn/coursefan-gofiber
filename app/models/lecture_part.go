package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type LecturePart struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	CourseId  uuid.UUID `json:"course_id" gorm:"type:CHAR(36);NOT NULL"`
	Course    *Course   `json:"course" gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Part      int       `json:"part"  gorm:"type:INTEGER UNSIGNED NOT NULL"`
	Name      string    `json:"title" gorm:"type:VARCHAR(50) NOT NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt null.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Lectures  []Lecture `json:"lectures" gorm:"foreignKey:LecturePartId;references:Id"`
}
