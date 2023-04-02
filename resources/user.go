package resources

import (
	"time"

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
