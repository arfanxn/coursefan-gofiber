package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id          uuid.UUID    `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	Name        string       `json:"name" gorm:"index;type:VARCHAR(50) NOT NULL"`
	Email       string       `json:"email" gorm:"uniqueIndex;type:VARCHAR(50) NOT NULL"`
	Password    string       `json:"password" gorm:"type:VARCHAR(256) NOT NULL"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt   null.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	UserProfile *UserProfile `json:"user_profile,omitempty"`
	UserSetting *UserSetting `json:"user_setting,omitempty"`

	ParticipatedCourses []Course `json:"participated_courses,omitempty" gorm:"many2many:course_user_role;"`

	// Avatar relationship
	Avatar *Media `json:"avatar" gorm:"-"`
}

func (User) TableName() string {
	return "users"
}

func (model User) FirstName() string {
	names := strings.Split(model.Name, " ")
	return names[0]
}

func (model User) LastName() string {
	names := strings.Split(model.Name, " ")
	if len(names) <= 1 {
		return ""
	}
	return strings.Join(names[1:len(names)-1], " ")
}
