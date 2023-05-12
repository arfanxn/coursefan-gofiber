package services

import (
	"net/url"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserSettingService struct {
	repository *repositories.UserSettingRepository
}

func NewUserSettingService(
	repository *repositories.UserSettingRepository,
) *UserSettingService {
	return &UserSettingService{
		repository: repository,
	}
}

// AllByCourse get users by course
func (service *UserSettingService) All(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.UserSetting], err error) {
	userMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	}
	usRess := sliceh.Map(userMdls, func(userMdl models.UserSetting) resources.UserSetting {
		usRes := resources.UserSetting{}
		usRes.FromModel(userMdl)
		return usRes
	})
	pagination.SetItems(usRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), int(input.Limit.Int64))
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *UserSettingService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.UserSetting, err error) {
	usMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	} else if len(usMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	data.FromModel(usMdls[0])
	return
}

// Update
func (service *UserSettingService) Update(c *fiber.Ctx, input requests.UserSettingUpdate) (
	usRes resources.UserSetting, err error) {
	usMdl := models.UserSetting{ // prepare for filtering
		UserId: uuid.MustParse(input.UserId),
		Key:    input.Key,
	}
	usMdl, err = service.repository.FindByModel(c, usMdl) // filter/find by model
	usMdl.Value = input.Value                             // assign the new value
	if errorh.IsGormErrRecordNotFound(err) {
		_, err = service.repository.Insert(c, &usMdl)
		if err != nil {
			return
		}
	} else if err != nil {
		return
	} else {
		_, err = service.repository.UpdateById(c, &usMdl)
		if err != nil {
			return
		}
	}
	usRes.FromModel(usMdl)
	return
}
