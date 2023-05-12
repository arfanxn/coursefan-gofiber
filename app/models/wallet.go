package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Wallet struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	OwnerId   uuid.UUID `json:"user_id" gorm:"type:CHAR(36);NOT NULL"`
	Owner     *User     `json:"user" gorm:"foreignKey:OwnerId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Balance   int64     `json:"balance"  gorm:"type:BIGINT UNSIGNED NOT NULL;default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
