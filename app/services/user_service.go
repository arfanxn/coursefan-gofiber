package services

import (
	"net/url"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserService struct {
	repository            *repositories.UserRepository
	mediaRepository       *repositories.MediaRepository
	userProfileRepository *repositories.UserProfileRepository
}

func NewUserService(
	repository *repositories.UserRepository,
	mediaRepository *repositories.MediaRepository,
	userProfileRepository *repositories.UserProfileRepository,
) *UserService {
	return &UserService{
		repository:            repository,
		mediaRepository:       mediaRepository,
		userProfileRepository: userProfileRepository,
	}
}

// AllByCourse get users by course
func (service *UserService) AllByCourse(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.User], err error) {
	userMdls, err := service.repository.AllByCourse(c, input)
	if err != nil {
		return
	}
	userRess := sliceh.Map(userMdls, func(userMdl models.User) resources.User {
		userRes := resources.User{}
		userRes.FromModel(userMdl)
		return userRes
	})
	pagination.SetItems(userRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), int(input.Limit.Int64))
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *UserService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.User, err error) {
	userMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	} else if len(userMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	data.FromModel(userMdls[0])
	return
}

// Update
func (service *UserService) Update(c *fiber.Ctx, input requests.UserUpdate) (
	userRes resources.User, err error) {
	syncronizer := synch.NewSyncronizer()

	syncronizer.WG().Add(3)
	go func() {
		defer syncronizer.WG().Done()
		userMdl, err := service.repository.FindById(c, input.Id)
		if errorh.IsGormErrRecordNotFound(err) {
			err = fiber.ErrNotFound
			syncronizer.Err(err)
			return
		} else if err != nil {
			syncronizer.Err(err)
			return
		}
		userMdl.Name = input.Name
		_, err = service.repository.UpdateById(c, &userMdl)
		if err != nil {
			syncronizer.Err(err)
			return
		}

		syncronizer.M().Lock()
		defer syncronizer.M().Unlock()
		userRes.FromModel(userMdl)
	}()

	go func() {
		defer syncronizer.WG().Done()
		upMdl, err := service.userProfileRepository.FindByUserId(c, input.Id)
		upMdl.UserId = uuid.MustParse(input.Id)
		upMdl.Headline = input.Headline
		upMdl.Biography = input.Biography
		upMdl.Language = input.Language
		upMdl.WebsiteUrl = input.WebsiteUrl
		upMdl.FacebookUrl = input.FacebookUrl
		upMdl.LinkedinUrl = input.LinkedinUrl
		upMdl.TwitterUrl = input.TwitterUrl
		upMdl.YoutubeUrl = input.YoutubeUrl

		if errorh.IsGormErrRecordNotFound(err) {
			_, err := service.userProfileRepository.Insert(c, &upMdl)
			if err != nil { // only return if error is happening
				syncronizer.Err(err)
				return
			}
		} else if err != nil {
			syncronizer.Err(err)
			return
		} else {
			_, err = service.userProfileRepository.UpdateById(c, &upMdl)
			if err != nil { // only return if error is happening
				syncronizer.Err(err)
				return
			}
		}

		syncronizer.M().Lock()
		defer syncronizer.M().Unlock()
		upRes := resources.UserProfile{}
		upRes.FromModel(upMdl)
		userRes.UserProfile = &upRes
	}()

	go func() {
		defer syncronizer.WG().Done()

		// if user's avatar not provided then return immediately
		if input.Avatar == nil {
			return
		}

		avatarMediaMdl, err := service.mediaRepository.FindByModel(c, models.Media{
			ModelType: reflecth.GetTypeName(models.User{}),
			ModelId:   uuid.MustParse(input.Id),
		})
		avatarMediaMdl.ModelType = reflecth.GetTypeName(models.User{})
		avatarMediaMdl.ModelId = uuid.MustParse(input.Id) // the user id
		avatarMediaMdl.CollectionName = null.StringFrom(enums.MediaCollectionNameUserAvatar)
		avatarMediaMdl.SetFileHeader(input.Avatar)
		if errorh.IsGormErrRecordNotFound(err) {
			_, err := service.mediaRepository.Insert(c, &avatarMediaMdl)
			if err != nil { // only return if error is happening
				syncronizer.Err(err)
				return
			}
		} else if err != nil {
			syncronizer.Err(err)
			return
		} else {
			_, err = service.mediaRepository.UpdateById(c, &avatarMediaMdl)
			if err != nil { // only return if error is happening
				syncronizer.Err(err)
				return
			}
		}

		syncronizer.M().Lock()
		defer syncronizer.M().Unlock()
		avatarMediaRes := resources.Media{}
		avatarMediaRes.FromModel(avatarMediaMdl)
		userRes.Avatar = &avatarMediaRes
	}()

	syncronizer.WG().Wait()
	
	if err != nil {
		return
	} else if err = syncronizer.Err(); err != nil {
		return
	}

	return
}
