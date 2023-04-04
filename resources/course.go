package resources

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Course struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   null.Time `json:"updated_at"`
}
