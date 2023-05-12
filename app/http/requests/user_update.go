package requests

import (
	"mime/multipart"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

type UserUpdate struct {
	// User Model
	Id   string `json:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required,min=2,max=50"`

	// User Profile Model
	Headline    null.String `json:"headline" validate:"max=50"`
	Biography   null.String `json:"biography" validate:"max=256"`
	Language    string      `json:"language" validate:"required,max=2,oneof=id en"`
	WebsiteUrl  null.String `json:"website_url" form:"website_url" validate:"max=256"`
	FacebookUrl null.String `json:"facebook_url" form:"facebook_url" validate:"max=256"`
	LinkedinUrl null.String `json:"linkedin_url" form:"linkedin_url" validate:"max=256"`
	TwitterUrl  null.String `json:"twitter_url" form:"twitter_url" validate:"max=256"`
	YoutubeUrl  null.String `json:"youtube_url" form:"youtube_url" validate:"max=256"`

	// User's Media Model
	Avatar *multipart.FileHeader `json:"avatar" form:"avatar" fhlidate:"max=10,mimes=image/jpeg"`
}

// FromContext fills input from the given context
func (input *UserUpdate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)

	if id := c.Params("user_id"); id != "" {
		input.Id = id
	}

	input.Avatar = ctxh.GetFileHeader(c, "avatar")

	return
}
