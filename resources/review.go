package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type Review struct {
	Id             string      `json:"id"`
	ReviewableType string      `json:"reviewable_type"`
	ReviewableId   string      `json:"reviewable_id"`
	ReviewerId     string      `json:"reviewer_id"`
	Reviewer       *User       `json:"reviewer,omitempty"`
	Rate           int         `json:"rate"`
	Title          null.String `json:"title"`
	Body           null.String `json:"body"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      null.Time   `json:"updated_at"`
}

func (resource *Review) FromModel(model models.Review) {
	resource.Id = model.Id.String()
	resource.ReviewableType = model.ReviewableType
	resource.ReviewableId = model.ReviewableId.String()
	resource.ReviewerId = model.ReviewerId.String()
	if model.Reviewer != nil {
		reviewerUserRes := User{}
		reviewerUserRes.FromModel(*model.Reviewer)
		resource.Reviewer = &reviewerUserRes
	}
	resource.Rate = model.Rate
	resource.Title = model.Title
	resource.Body = model.Body
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
