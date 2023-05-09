package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt null.Time `json:"updated_at"`

	Avatar Media `json:"avatar,omitempty"`
}

func (resource *User) FromModel(model models.User) {
	resource.Id = model.Id.String()
	resource.Name = model.Name
	resource.Email = model.Email
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt

	if model.Avatar != nil {
		avatarMediaRes := Media{}
		avatarMediaRes.FromModel(*model.Avatar)
		resource.Avatar = avatarMediaRes
	}
}
