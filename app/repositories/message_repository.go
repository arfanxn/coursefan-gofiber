package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

// NewMessageRepository instantiates a new MessageRepository
func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// All returns all messages in the database
func (repository *MessageRepository) All(c *fiber.Ctx) (messages []models.Message, err error) {
	err = repository.db.Find(&messages).Error
	return
}

// Find finds model by id
func (repository *MessageRepository) Find(c *fiber.Ctx, id string) (message models.Message, err error) {
	err = repository.db.Where("id = ?", id).First(&message).Error
	return
}

// Insert inserts models into the database
func (repository *MessageRepository) Insert(c *fiber.Ctx, messages ...*models.Message) (int64, error) {
	for _, message := range messages {
		if message.Id == uuid.Nil {
			message.Id = uuid.New()
		}
		message.CreatedAt = time.Now()
	}
	result := repository.db.Create(messages)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *MessageRepository) UpdateById(c *fiber.Ctx, message *models.Message) (int64, error) {
	// refresh model updated at
	message.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(message).Updates(message)
	return result.RowsAffected, result.Error
}
