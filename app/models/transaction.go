package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Transaction struct {
	Id                  uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	TransactionableType string    `json:"transactionable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	TransactionableId   uuid.UUID `json:"transactionable_id" gorm:"type:CHAR(36) NOT NULL"`
	SenderId            uuid.UUID `json:"sender_id" gorm:"type:CHAR(36);NOT NULL"`
	Sender              Wallet    `json:"sender" gorm:"foreignKey:SenderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReceiverId          uuid.UUID `json:"receiver_id" gorm:"type:CHAR(36);NOT NULL"`
	Receiver            Wallet    `json:"receiver" gorm:"foreignKey:ReceiverId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount              int64     `json:"amount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Rate                int64     `json:"rate"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Discount            int64     `json:"discount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Total               int64     `json:"total"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	CancelledAt         null.Time `json:"cancelled_at" gorm:"type:DATETIME(3)"`
	ChargebackedAt      null.Time `json:"chargebacked_at" gorm:"type:DATETIME(3)"`
	ExpiredAt           null.Time `json:"expired_at" gorm:"type:DATETIME(3)"`
	FailedAt            null.Time `json:"failed_at" gorm:"type:DATETIME(3)"`
	RefundedAt          null.Time `json:"refunded_at" gorm:"type:DATETIME(3)"`
	SettledAt           null.Time `json:"settled_at" gorm:"type:DATETIME(3)"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt           null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}