package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserSetting struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	User      *User     `json:"user,omitempty"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`
}

func (resource *UserSetting) FromModel(model models.UserSetting) {
	resource.Id = model.Id.String()
	resource.UserId = model.UserId.String()
	if model.User.Id != uuid.Nil {
		userRes := User{}
		userRes.FromModel(*model.User)
		resource.User = &userRes
	}
	resource.Key = model.Key
	resource.Value = model.Value
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
