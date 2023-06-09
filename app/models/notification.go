package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Notification struct {
	Id         uuid.UUID     `json:"id" gorm:"primaryKey;type:char(36)"`
	SenderId   uuid.UUID     `json:"sender_id" gorm:"type:CHAR(36);NOT NULL"`
	Sender     *User         `json:"sender" gorm:"foreignKey:SenderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReceiverId uuid.UUID     `json:"receiver_id" gorm:"type:CHAR(36);NOT NULL"`
	Receiver   *User         `json:"receiver" gorm:"foreignKey:ReceiverId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ObjectType null.String   `json:"object_type" gorm:"index;type:VARCHAR(25)"`
	ObjectId   uuid.NullUUID `json:"object_id" gorm:"type:VARCHAR(36)"`
	Object     any           `json:"object" gorm:"-"`
	Title      string        `json:"title" gorm:"type:VARCHAR(50) NOT NULL"`
	Body       null.String   `json:"body" gorm:"type:TEXT"`
	Type       null.String   `json:"type"  gorm:"type:VARCHAR(25)"`
	ReadedAt   null.Time     `json:"readed_at" gorm:"type:DATETIME(3)"`
	CreatedAt  time.Time     `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt  null.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Notification) TableName() string {
	return "notifications"
}
