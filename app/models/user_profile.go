package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserProfile struct {
	Id          uuid.UUID   `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	UserId      uuid.UUID   `json:"user_id" gorm:"type:CHAR(36);NOT NULL"`
	User        User        `json:"user" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Headline    null.String `json:"headline" gorm:"type:VARCHAR(50)"`
	Biography   null.String `json:"biography" gorm:"type:VARCHAR(256)"`
	Language    string      `json:"language" gorm:"type:CHAR(2) NOT NULL"`
	WebsiteUrl  null.String `json:"website_url" gorm:"type:VARCHAR(256)"`
	FacebookUrl null.String `json:"facebook_url" gorm:"type:VARCHAR(256)"`
	LinkedinUrl null.String `json:"linkedin_url" gorm:"type:VARCHAR(256)"`
	TwitterUrl  null.String `json:"twitter_url" gorm:"type:VARCHAR(256)"`
	YoutubeUrl  null.String `json:"youtube_url" gorm:"type:VARCHAR(256)"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt   null.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}
