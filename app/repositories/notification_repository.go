package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

// NewNotificationRepository instantiates a new NotificationRepository
func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// All returns all notifications in the database
func (repository *NotificationRepository) All(c *fiber.Ctx) (notifications []models.Notification, err error) {
	err = repository.db.Find(&notifications).Error
	return
}

// AllByAuthUser returns all notifications by current logged in user
func (repository *NotificationRepository) AllByAuthUser(c *fiber.Ctx, query requests.Query) (
	notifications []models.Notification, err error) {
	err = gormh.BuildFromRequestQuery(repository.db, models.Review{}, query).
		Where(models.Notification{}.TableName()+".receiver_id = ?", ctxh.MustGetUser(c).Id.String()).
		Distinct().Find(&notifications).Error

	return
}

// Find finds model by id
func (repository *NotificationRepository) FindById(c *fiber.Ctx, id string) (notification models.Notification, err error) {
	err = repository.db.Where("id = ?", id).First(&notification).Error
	return
}

// Insert inserts models into the database
func (repository *NotificationRepository) Insert(c *fiber.Ctx, notifications ...*models.Notification) (int64, error) {
	for _, notification := range notifications {
		if notification.Id == uuid.Nil {
			notification.Id = uuid.New()
		}
		notification.CreatedAt = time.Now()
	}
	result := repository.db.Create(notifications)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *NotificationRepository) UpdateById(c *fiber.Ctx, notification *models.Notification) (int64, error) {
	// refresh model updated at
	notification.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(notification).Updates(notification)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *NotificationRepository) DeleteByIds(c *fiber.Ctx, notifications ...*models.Notification) (int64, error) {
	result := repository.db.Delete(notifications)
	return result.RowsAffected, result.Error
}
