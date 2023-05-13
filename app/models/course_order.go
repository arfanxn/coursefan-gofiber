package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type CourseOrder struct {
	Id             uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	UserId         uuid.UUID `json:"user_id" gorm:"type:CHAR(36);NOT NULL"`
	User           *User     `json:"user" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CourseId       uuid.UUID `json:"course_id" gorm:"type:CHAR(36);NOT NULL"`
	Course         *Course   `json:"course" gorm:"foreignKey:CourseId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount         float64   `json:"amount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Rate           float64   `json:"rate"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Discount       float64   `json:"discount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Total          float64   `json:"total"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	CancelledAt    null.Time `json:"cancelled_at" gorm:"type:DATETIME(3)"`
	ChargebackedAt null.Time `json:"chargebacked_at" gorm:"type:DATETIME(3)"`
	ExpiredAt      null.Time `json:"expired_at" gorm:"type:DATETIME(3)"`
	FailedAt       null.Time `json:"failed_at" gorm:"type:DATETIME(3)"`
	RefundedAt     null.Time `json:"refunded_at" gorm:"type:DATETIME(3)"`
	SettledAt      null.Time `json:"settled_at" gorm:"type:DATETIME(3)"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt      null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (CourseOrder) TableName() string {
	return "course_orders"
}
