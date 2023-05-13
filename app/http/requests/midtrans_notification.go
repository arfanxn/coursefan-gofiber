package requests

import (
	"github.com/gofiber/fiber/v2"
)

type MidtransNotification struct {
	VANumbers         []MidtransNotificationVANumbers
	TransactionId     string  `json:"transaction_id,omitempty" form:"transaction_id"`
	TransactionTime   string  `json:"transaction_time,omitempty" form:"transaction_time"`
	TransactionStatus string  `json:"transaction_status,omitempty" form:"transaction_status"`
	StatusMessage     string  `json:"status_message,omitempty" form:"status_message"`
	StatusCode        int     `json:"status_code,string,omitempty" form:"status_code"`
	SignatureKey      string  `json:"signature_key,omitempty" form:"signature_key"`
	PaymentType       string  `json:"payment_type,omitempty" form:"payment_type"`
	PaymentAmounts    []any   `json:"payment_amounts,omitempty" form:"payment_amounts"`
	OrderId           string  `json:"order_id,omitempty" form:"order_id"`
	MerchantId        string  `json:"merchant_id,omitempty" form:"merchant_id"`
	GrossAmount       float64 `json:"gross_amount,string,omitempty" form:"gross_amount"`
	FraudStatus       string  `json:"fraud_status,omitempty" form:"fraud_status"`
	ExpiryTime        string  `json:"expiry_time,omitempty" form:"expiry_time"`
	Currency          string  `json:"currency,omitempty" form:"currency"`
}

type MidtransNotificationVANumbers struct {
	VANumber string `json:"va_number,omitempty" form:"va_number"`
	Bank     string `json:"bank,omitempty" form:"bank"`
}

// FromContext fills input from the given context
func (input *MidtransNotification) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	return
}
