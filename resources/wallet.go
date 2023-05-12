package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type Wallet struct {
	Id        string    `json:"id"`
	OwnerId   string    `json:"owner_id"`
	Owner     *User     `json:"owner,omitempty"`
	Balance   int64     `json:"balance" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func (resource *Wallet) FromModel(model models.Wallet) {
	resource.Id = model.Id.String()
	resource.OwnerId = model.OwnerId.String()
	if model.Owner != nil {
		userRes := User{}
		userRes.FromModel(*model.Owner)
		resource.Owner = &userRes
	}
	resource.Balance = model.Balance
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
