package services

import (
	"net/url"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ReviewService struct {
	repository *repositories.ReviewRepository
}

func NewReviewService(repository *repositories.ReviewRepository) *ReviewService {
	return &ReviewService{
		repository: repository,
	}
}

// AllByCourse get lecture parts by course
func (service *ReviewService) AllByCourse(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.Review], err error) {
	reviewMdls, err := service.repository.AllByCourse(c, input)
	if err != nil {
		return
	}
	reviewRess := sliceh.Map(reviewMdls, func(reviewMdl models.Review) resources.Review {
		reviewRes := resources.Review{}
		reviewRes.FromModel(reviewMdl)
		return reviewRes
	})
	pagination.SetItems(reviewRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), int(input.Limit.Int64))
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *ReviewService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.Review, err error) {
	reviewMdls, err := service.repository.AllByCourse(c, input)
	if err != nil {
		return
	} else if len(reviewMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	data.FromModel(reviewMdls[0])
	return
}

// Create
func (service *ReviewService) Create(c *fiber.Ctx, input requests.ReviewCreate) (
	reviewRes resources.Review, err error) {
	reviewMdl := models.Review{}
	reviewMdl.ReviewableType = input.ReviewableType
	reviewMdl.ReviewableId = uuid.MustParse(input.ReviewableId)
	reviewMdl.ReviewerId = uuid.MustParse(input.ReviewerId)
	reviewMdl.Rate = input.Rate
	reviewMdl.Title = input.Title
	reviewMdl.Body = input.Body
	_, err = service.repository.Insert(c, &reviewMdl)
	if err != nil {
		return
	}
	reviewRes.FromModel(reviewMdl)
	return
}

// Create By Course
func (service *ReviewService) CreateByCourse(c *fiber.Ctx, input requests.ReviewCreate) (
	reviewRes resources.Review, err error) {
	reviewMdl := models.Review{}
	reviewMdl.ReviewableType = reflecth.GetTypeName(models.Course{})
	reviewMdl.ReviewableId = uuid.MustParse(input.ReviewableId)
	reviewMdl.ReviewerId = uuid.MustParse(input.ReviewerId)
	reviewMdl.Rate = input.Rate
	reviewMdl.Title = input.Title
	reviewMdl.Body = input.Body
	_, err = service.repository.Insert(c, &reviewMdl)
	if err != nil {
		return
	}
	reviewRes.FromModel(reviewMdl)
	return
}

// Update
func (service *ReviewService) Update(c *fiber.Ctx, input requests.ReviewUpdate) (
	reviewRes resources.Review, err error) {
	reviewMdl, err := service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	reviewMdl.Rate = input.Rate
	reviewMdl.Title = input.Title
	reviewMdl.Body = input.Body

	spew.Dump(reviewMdl)

	_, err = service.repository.UpdateById(c, &reviewMdl)
	if err != nil {
		return
	}
	reviewRes.FromModel(reviewMdl)
	return
}

// Delete
func (service *ReviewService) Delete(c *fiber.Ctx, input requests.ReviewDelete) (err error) {
	var reviewMdl models.Review
	reviewMdl, err = service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	_, err = service.repository.DeleteByIds(c, &reviewMdl)
	if err != nil {
		return
	}
	return
}
