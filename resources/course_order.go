package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type CourseOrder struct {
	Id             string    `json:"id"`
	UserId         string    `json:"user_id"`
	User           *User     `json:"user"`
	CourseId       string    `json:"course_id"`
	Course         *Course   `json:"course"`
	Amount         float64   `json:"amount"`
	Rate           float64   `json:"rate"`
	Discount       float64   `json:"discount"`
	Total          float64   `json:"total"`
	CancelledAt    null.Time `json:"cancelled_at"`
	ChargebackedAt null.Time `json:"chargebacked_at"`
	ExpiredAt      null.Time `json:"expired_at"`
	FailedAt       null.Time `json:"failed_at"`
	RefundedAt     null.Time `json:"refunded_at"`
	SettledAt      null.Time `json:"settled_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      null.Time `json:"updated_at"`
}

func (resource *CourseOrder) FromModel(model models.CourseOrder) {
	resource.Id = model.Id.String()
	resource.UserId = model.UserId.String()
	if model.User != nil {
		userRes := User{}
		userRes.FromModel(*model.User)
		resource.User = &userRes
	}
	resource.CourseId = model.CourseId.String()
	if model.Course != nil {
		courseRes := Course{}
		courseRes.FromModel(*model.Course)
		resource.Course = &courseRes
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
