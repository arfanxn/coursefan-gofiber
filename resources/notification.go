package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type Notification struct {
	Id         string      `json:"id"`
	SenderId   string      `json:"sender_id"`
	Sender     *User       `json:"sender,omitempty"`
	ReceiverId string      `json:"receiver_id"`
	Receiver   *User       `json:"receiver,omitempty"`
	ObjectType null.String `json:"object_type"`
	ObjectId   null.String `json:"object_id"`
	Object     any         `json:"object,omitempty"`
	Title      string      `json:"title"`
	Body       null.String `json:"body"`
	Type       null.String `json:"type"`
	ReadedAt   null.Time   `json:"readed_at"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  null.Time   `json:"updated_at"`
}

func (resource *Notification) FromModel(model models.Notification) {
	resource.Id = model.Id.String()
	resource.SenderId = model.SenderId.String()
	if model.Sender != nil {
		senderUserRes := User{}
		senderUserRes.FromModel(*model.Sender)
		resource.Sender = &senderUserRes
	}
	resource.ReceiverId = model.ReceiverId.String()
	if model.Receiver != nil {
		receiverUserRes := User{}
		receiverUserRes.FromModel(*model.Receiver)
		resource.Receiver = &receiverUserRes
	}
	resource.ObjectId = null.NewString(model.ObjectId.UUID.String(), model.ObjectId.Valid)
	resource.ObjectType = model.ObjectType
	resource.Object = model.Object
	resource.Title = model.Title
	resource.Body = model.Body
	resource.Type = model.Type
	resource.ReadedAt = model.ReadedAt
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = model.UpdatedAt
}
