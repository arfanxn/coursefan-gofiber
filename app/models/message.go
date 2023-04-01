package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Message struct {
	Id         uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	SenderId   uuid.UUID `json:"sender_id" gorm:"type:CHAR(36);NOT NULL"`
	Sender     User      `json:"sender" gorm:"foreignKey:SenderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReceiverId uuid.UUID `json:"receiver_id" gorm:"type:CHAR(36);NOT NULL"`
	Receiver   User      `json:"receiver" gorm:"foreignKey:ReceiverId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Body       string    `json:"body" gorm:"type:TEXT;NOT NULL"`
	ReadedAt   null.Time `json:"readed_at" gorm:"type:DATETIME(3)"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt  null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
