package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Discussion struct {
	Id                  string       `json:"id"`
	DiscussableType     string       `json:"discussable_type"`
	DiscussableId       string       `json:"discussable_id"`
	DiscusserId         string       `json:"discusser_id"`
	Discusser           User         `json:"discusser,omitempty"`
	DiscussionRepliedId null.String  `json:"discussion_replied_id"`
	DiscussionReplied   *Discussion  `json:"discussion_replied,omitempty"`
	DiscussionReplies   []Discussion `json:"discussion_replies,omitempty"`
	Title               string       `json:"title"`
	Body                null.String  `json:"body"`
	Upvote              int          `json:"upvote"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           null.Time    `json:"updated_at"`
}

func (resource *Discussion) FromModel(model models.Discussion) {
	resource.Id = model.Id.String()
	resource.DiscussableType = model.DiscussableType
	resource.DiscussableId = model.DiscussableId.String()
	resource.DiscusserId = model.DiscusserId.String()
	if model.Discusser.Id != uuid.Nil {
		discusserUserRes := User{}
		discusserUserRes.FromModel(model.Discusser)
		resource.Discusser = discusserUserRes
	}
	resource.DiscussionRepliedId = null.NewString(
		model.DiscussionRepliedId.UUID.String(),
		model.DiscussionRepliedId.Valid)
	if model.DiscussionReplied != nil {
		discussionRes := Discussion{}
		discussionRes.FromModel(*model.DiscussionReplied)
		resource.DiscussionReplied = &discussionRes
	}
	for _, discussionReplyMdl := range model.DiscussionReplies {
		discussionReplyRes := Discussion{}
		discussionReplyRes.FromModel(discussionReplyMdl)
		resource.DiscussionReplies = append(resource.DiscussionReplies, discussionReplyRes)
	}
	resource.Title = model.Title
	resource.Body = model.Body
	resource.Upvote = model.Upvote
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
