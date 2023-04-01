package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id                  uuid.UUID    `json:"id" gorm:"primaryKey;type:char(36)"`
	TransactionableType string       `json:"transactionable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	TransactionableId   uuid.UUID    `json:"transactionable_id" gorm:"type:CHAR(36) NOT NULL"`
	TransactorId        uuid.UUID    `json:"transactor_id" gorm:"type:CHAR(36);NOT NULL"`
	Transactor          User         `json:"transactor" gorm:"foreignKey:TransactorId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactedId        uuid.UUID    `json:"transacted_id" gorm:"type:CHAR(36);NOT NULL"`
	Transacted          User         `json:"transacted" gorm:"foreignKey:TransactedId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount              int64        `json:"amount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Rate                int64        `json:"rate"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Discount            int64        `json:"discount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Total               int64        `json:"total"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	CancelledAt         sql.NullTime `json:"cancelled_at" gorm:"type:DATETIME(3)"`
	ChargebackedAt      sql.NullTime `json:"chargebacked_at" gorm:"type:DATETIME(3)"`
	ExpiredAt           sql.NullTime `json:"expired_at" gorm:"type:DATETIME(3)"`
	FailedAt            sql.NullTime `json:"failed_at" gorm:"type:DATETIME(3)"`
	RefundedAt          sql.NullTime `json:"refunded_at" gorm:"type:DATETIME(3)"`
	SettledAt           sql.NullTime `json:"settled_at" gorm:"type:DATETIME(3)"`
	CreatedAt           time.Time    `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt           sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}
