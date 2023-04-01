package models

import (
	"github.com/google/uuid"
)

type TransactionWallet struct {
	TransactionId uuid.UUID   `json:"transaction_id" gorm:"type:CHAR(36);NOT NULL"`
	Transaction   Transaction `json:"transaction" gorm:"foreignKey:TransactionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SenderId      uuid.UUID   `json:"sender_id" gorm:"type:CHAR(36);NOT NULL"`
	Sender        Wallet      `json:"sender" gorm:"foreignKey:SenderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReceiverId    uuid.UUID   `json:"receiver_id" gorm:"type:CHAR(36);NOT NULL"`
	Receiver      Wallet      `json:"receiver" gorm:"foreignKey:ReceiverId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (TransactionWallet) TableName() string {
	return "transaction_wallet"
}
