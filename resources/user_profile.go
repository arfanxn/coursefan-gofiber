package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserProfile struct {
	Id          string      `json:"id"`
	UserId      string      `json:"user_id"`
	User        *User       `json:"user,omitempty"`
	Headline    null.String `json:"headline"`
	Biography   null.String `json:"biography"`
	Language    string      `json:"language"`
	WebsiteUrl  null.String `json:"website_url"`
	FacebookUrl null.String `json:"facebook_url"`
	LinkedinUrl null.String `json:"linkedin_url"`
	TwitterUrl  null.String `json:"twitter_url"`
	YoutubeUrl  null.String `json:"youtube_url"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   null.Time   `json:"updated_at"`
}

func (resource *UserProfile) FromModel(model models.UserProfile) {
	resource.Id = model.Id.String()
	resource.UserId = model.UserId.String()
	if (model.User != nil) && (model.User.Id != uuid.Nil) {
		userRes := User{}
		userRes.FromModel(*model.User)
		resource.User = &userRes
	}
	resource.Headline = model.Headline
	resource.Biography = model.Biography
	resource.Language = model.Language
	resource.WebsiteUrl = model.WebsiteUrl
	resource.FacebookUrl = model.FacebookUrl
	resource.LinkedinUrl = model.LinkedinUrl
	resource.TwitterUrl = model.TwitterUrl
	resource.YoutubeUrl = model.YoutubeUrl
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
