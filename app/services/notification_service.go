package services

import (
	"net/url"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/nullh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type NotificationService struct {
	repository *repositories.NotificationRepository
}

func NewNotificationService(repository *repositories.NotificationRepository) *NotificationService {
	return &NotificationService{
		repository: repository,
	}
}

// AllByAuthUser get discussions by lecture
func (service *NotificationService) AllByAuthUser(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.Notification], err error) {
	notificationMdls, err := service.repository.AllByAuthUser(c, input)
	if err != nil {
		return
	}
	notificationRess := sliceh.Map(notificationMdls, func(notificationMdl models.Notification) resources.Notification {
		notificationRes := resources.Notification{}
		notificationRes.FromModel(notificationMdl)
		return notificationRes
	})
	pagination.SetItems(notificationRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), int(input.Limit.Int64))
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *NotificationService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.Notification, err error) {
	notificationMdls, err := service.repository.AllByAuthUser(c, input)
	notificationMdl := sliceh.FirstOrNil(notificationMdls)
	if err != nil {
		return
	} else if notificationMdl == nil {
		err = fiber.ErrNotFound
		return
	} else if sliceh.NotContains([]uuid.UUID{notificationMdl.SenderId, notificationMdl.ReceiverId}, ctxh.MustGetUser(c).Id) {
		err = fiber.ErrForbidden
		return
	}
	data.FromModel(*notificationMdl)
	return
}

// Create
func (service *NotificationService) Create(c *fiber.Ctx, input requests.NotificationCreate) (
	notificationRes resources.Notification, err error) {
	notificationMdl := models.Notification{}
	notificationMdl.SenderId = uuid.MustParse(input.SenderId)
	notificationMdl.ReceiverId = uuid.MustParse(input.ReceiverId)
	notificationMdl.ObjectType = input.ObjectType
	notificationMdl.ObjectId = nullh.NullStringToNullUUID(input.ObjectId)
	notificationMdl.Title = input.Title
	notificationMdl.Body = input.Body
	notificationMdl.Type = input.Type
	_, err = service.repository.Insert(c, &notificationMdl)
	if err != nil {
		return
	}
	notificationRes.FromModel(notificationMdl)
	return
}

// Update
func (service *NotificationService) Update(c *fiber.Ctx, input requests.NotificationUpdate) (
	notificationRes resources.Notification, err error) {
	notificationMdl, err := service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	} else if notificationMdl.SenderId != ctxh.MustGetUser(c).Id {
		err = fiber.ErrForbidden
		return
	}

	notificationMdl.ObjectType = input.ObjectType
	notificationMdl.ObjectId = nullh.NullStringToNullUUID(input.ObjectId)
	notificationMdl.Title = input.Title
	notificationMdl.Body = input.Body
	notificationMdl.Type = input.Type

	_, err = service.repository.UpdateById(c, &notificationMdl)
	if err != nil {
		return
	}
	notificationRes.FromModel(notificationMdl)
	return
}

// Mark marks notification as read or unread by the specified boolean argument
func (service *NotificationService) Mark(c *fiber.Ctx, input requests.NotificationId, markRead bool) (
	notificationRes resources.Notification, err error) {
	notificationMdl, err := service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	} else if notificationMdl.ReceiverId != ctxh.MustGetUser(c).Id {
		err = fiber.ErrForbidden
		return
	}

	if markRead {
		notificationMdl.ReadedAt = null.NewTime(time.Now(), true)
	} else if !markRead {
		notificationMdl.ReadedAt = null.NewTime(time.Time{}, false)
	}
	_, err = service.repository.UpdateById(c, &notificationMdl)
	if err != nil {
		return
	}
	notificationRes.FromModel(notificationMdl)
	return
}

// Delete
func (service *NotificationService) Delete(c *fiber.Ctx, input requests.NotificationId) (err error) {
	var notificationMdl models.Notification
	notificationMdl, err = service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	} else if notificationMdl.SenderId != ctxh.MustGetUser(c).Id {
		err = fiber.ErrForbidden
		return
	}
	_, err = service.repository.DeleteByIds(c, &notificationMdl)
	if err != nil {
		return
	}
	return
}
