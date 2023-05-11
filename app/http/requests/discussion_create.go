package requests

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
)

type DiscussionCreate struct {
	DiscussableType     string      `json:"discussable_type" form:"discussable_type" validate:"oneof=Lecture"`
	DiscussableId       string      `json:"discussable_id" form:"discussable_id"  validate:"required,uuid"`
	DiscussionRepliedId null.String `json:"discussion_replied_id" form:"discussion_replied_id"`
	Title               string      `json:"title" validate:"required,min=10"`
	Body                null.String `json:"body"`
}

// FromContext fills input from the given context
func (input *DiscussionCreate) FromContext(c *fiber.Ctx) (err error) {
	err = c.BodyParser(input)
	return
}
