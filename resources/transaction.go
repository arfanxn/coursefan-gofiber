package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type Transaction struct {
	Id                  string    `json:"id" gorm:"primaryKey;type:char(36)"`
	TransactionableType string    `json:"transactionable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	TransactionableId   string    `json:"transactionable_id" gorm:"type:CHAR(36) NOT NULL"`
	SenderId            string    `json:"sender_id" gorm:"type:CHAR(36);NOT NULL"`
	Sender              *Wallet   `json:"sender" gorm:"foreignKey:SenderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReceiverId          string    `json:"receiver_id" gorm:"type:CHAR(36);NOT NULL"`
	Receiver            *Wallet   `json:"receiver" gorm:"foreignKey:ReceiverId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount              float64   `json:"amount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Rate                float64   `json:"rate"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Discount            float64   `json:"discount"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	Total               float64   `json:"total"  gorm:"type:BIGINT UNSIGNED NOT NULL"`
	CancelledAt         null.Time `json:"cancelled_at" gorm:"type:DATETIME(3)"`
	ChargebackedAt      null.Time `json:"chargebacked_at" gorm:"type:DATETIME(3)"`
	ExpiredAt           null.Time `json:"expired_at" gorm:"type:DATETIME(3)"`
	FailedAt            null.Time `json:"failed_at" gorm:"type:DATETIME(3)"`
	RefundedAt          null.Time `json:"refunded_at" gorm:"type:DATETIME(3)"`
	SettledAt           null.Time `json:"settled_at" gorm:"type:DATETIME(3)"`
	CreatedAt           time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt           null.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (resource *Transaction) FromModel(model models.Transaction) {
	resource.Id = model.Id.String()
	resource.TransactionableType = model.TransactionableType
	resource.TransactionableId = model.TransactionableId.String()
	resource.SenderId = model.SenderId.String()
	if model.Sender != nil {
		senderWalletRes := Wallet{}
		senderWalletRes.FromModel(*model.Sender)
		resource.Sender = &senderWalletRes
	}
	resource.ReceiverId = model.ReceiverId.String()
	if model.Receiver != nil {
		senderWalletRes := Wallet{}
		senderWalletRes.FromModel(*model.Receiver)
		resource.Receiver = &senderWalletRes
	}
	resource.Amount = model.Amount
	resource.Rate = model.Rate
	resource.Discount = model.Discount
	resource.Total = model.Total

	resource.CancelledAt = model.CancelledAt
	resource.ChargebackedAt = model.ChargebackedAt
	resource.ExpiredAt = model.ExpiredAt
	resource.FailedAt = model.FailedAt
	resource.RefundedAt = model.RefundedAt
	resource.SettledAt = model.SettledAt
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
