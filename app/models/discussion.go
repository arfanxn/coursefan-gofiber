package models

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Discussion struct {
	Id                  uuid.UUID     `json:"id" gorm:"primaryKey;type:char(36)"`
	DiscussableType     string        `json:"discussable_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	DiscussableId       uuid.UUID     `json:"discussable_id" gorm:"type:VARCHAR(36) NOT NULL"`
	DiscusserId         uuid.UUID     `json:"discusser_id" gorm:"type:CHAR(36);NOT NULL"`
	Discusser           User          `json:"discusser" gorm:"foreignKey:DiscusserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DiscussionRepliedId uuid.NullUUID `json:"discussion_replied_id" gorm:"type:CHAR(36);comment:this column used for store the parent discussion id (replied discussion id), this can be null, if null it means the discussion is the parent (main discussion)"`
	DiscussionReplied   *Discussion   `json:"discussion_replied" gorm:"foreignKey:DiscussionRepliedId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DiscussionReplies   []Discussion  `json:"discussion_replies" gorm:"foreignKey:DiscussionRepliedId"`
	Title               string        `json:"title" gorm:"type:VARCHAR(50) NOT NULL"`
	Body                null.String   `json:"body" gorm:"type:TEXT"`
	Upvote              int           `json:"upvote"  gorm:"index;type:INTEGER UNSIGNED NOT NULL;default:0"`
	CreatedAt           time.Time     `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	UpdatedAt           null.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Discussion) TableName() string {
	return "discussions"
}
