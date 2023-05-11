package services

import (
	"net/url"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/nullh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DiscussionService struct {
	repository *repositories.DiscussionRepository
}

func NewDiscussionService(repository *repositories.DiscussionRepository) *DiscussionService {
	return &DiscussionService{
		repository: repository,
	}
}

// AllByLecture get discussions by lecture
func (service *DiscussionService) AllByLecture(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.Discussion], err error) {
	discussionMdls, err := service.repository.AllByLecture(c, input)
	if err != nil {
		return
	}
	discussionRess := sliceh.Map(discussionMdls, func(discussionMdl models.Discussion) resources.Discussion {
		discussionRes := resources.Discussion{}
		discussionRes.FromModel(discussionMdl)
		return discussionRes
	})
	pagination.SetItems(discussionRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), int(input.Limit.Int64))
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *DiscussionService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.Discussion, err error) {
	discussionMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	} else if len(discussionMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	data.FromModel(discussionMdls[0])
	return
}

// Create
func (service *DiscussionService) Create(c *fiber.Ctx, input requests.DiscussionCreate) (
	discussionRes resources.Discussion, err error) {
	discussionMdl := models.Discussion{}
	discussionMdl.DiscussableType = input.DiscussableType
	discussionMdl.DiscussableId = uuid.MustParse(input.DiscussableId)
	discussionMdl.DiscussionRepliedId = nullh.NullStringToNullUUID(input.DiscussionRepliedId)
	discussionMdl.DiscusserId = ctxh.MustGetUser(c).Id
	discussionMdl.Title = input.Title
	discussionMdl.Body = input.Body
	_, err = service.repository.Insert(c, &discussionMdl)
	if err != nil {
		return
	}
	discussionRes.FromModel(discussionMdl)
	return
}

// Create By Lecture
func (service *DiscussionService) CreateByLecture(c *fiber.Ctx, input requests.DiscussionCreate) (
	discussionRes resources.Discussion, err error) {
	discussionMdl := models.Discussion{}
	discussionMdl.DiscussableType = reflecth.GetTypeName(models.Lecture{})
	discussionMdl.DiscussableId = uuid.MustParse(input.DiscussableId)
	discussionMdl.DiscussionRepliedId = nullh.NullStringToNullUUID(input.DiscussionRepliedId)
	discussionMdl.DiscusserId = ctxh.MustGetUser(c).Id
	discussionMdl.Title = input.Title
	discussionMdl.Body = input.Body
	_, err = service.repository.Insert(c, &discussionMdl)
	if err != nil {
		return
	}
	discussionRes.FromModel(discussionMdl)
	return
}

// Update
func (service *DiscussionService) Update(c *fiber.Ctx, input requests.DiscussionUpdate) (
	discussionRes resources.Discussion, err error) {
	discussionMdl, err := service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	discussionMdl.Title = input.Title
	discussionMdl.Body = input.Body

	_, err = service.repository.UpdateById(c, &discussionMdl)
	if err != nil {
		return
	}
	discussionRes.FromModel(discussionMdl)
	return
}

// Upvote
func (service *DiscussionService) Upvote(c *fiber.Ctx, input requests.DiscussionId) (
	discussionRes resources.Discussion, err error) {
	discussionMdl, err := service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	discussionMdl.Upvote = discussionMdl.Upvote + 1

	_, err = service.repository.UpdateById(c, &discussionMdl)
	if err != nil {
		return
	}
	discussionRes.FromModel(discussionMdl)
	return
}

// Delete
func (service *DiscussionService) Delete(c *fiber.Ctx, input requests.DiscussionId) (err error) {
	var discussionMdl models.Discussion
	discussionMdl, err = service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	_, err = service.repository.DeleteByIds(c, &discussionMdl)
	if err != nil {
		return
	}
	return
}
