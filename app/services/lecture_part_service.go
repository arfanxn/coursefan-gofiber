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

type LecturePartService struct {
	repository *repositories.LecturePartRepository
}

func NewLecturePartService(repository *repositories.LecturePartRepository) *LecturePartService {
	return &LecturePartService{
		repository: repository,
	}
}

// AllByCourse get lecture parts by course
func (service *LecturePartService) AllByCourse(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.LecturePart], err error) {
	lecturePartMdls, err := service.repository.AllByCourse(c, input)
	if err != nil {
		return
	}
	lecturePartRess := sliceh.Map(lecturePartMdls, func(lpMdl models.LecturePart) resources.LecturePart {
		lpRes := resources.LecturePart{}
		lpRes.FromModel(lpMdl)
		return lpRes
	})
	pagination.SetItems(lecturePartRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), input.Limit)
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *LecturePartService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.LecturePart, err error) {
	lecturePartMdls, err := service.repository.AllByCourse(c, input)
	if err != nil {
		return
	} else if len(lecturePartMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	data.FromModel(lecturePartMdls[0])
	return
}

// Create
func (service *LecturePartService) Create(c *fiber.Ctx, input requests.LecturePartCreate) (
	data resources.LecturePart, err error) {
	lecturePartMdl := models.LecturePart{}
	lecturePartMdl.Part = input.Part
	lecturePartMdl.Name = input.Name
	_, err = service.repository.Insert(c, &lecturePartMdl)
	if err != nil {
		return
	}
	data.FromModel(lecturePartMdl)
	return
}

// Update
func (service *LecturePartService) Update(c *fiber.Ctx, input requests.LecturePartUpdate) (
	lpRes resources.LecturePart, err error) {
	lpMdl, err := service.repository.FindByModel(c, models.LecturePart{
		Id:       uuid.MustParse(input.Id),
		CourseId: uuid.MustParse(input.CourseId),
	})
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	lpMdl.Name = input.Name
	lpMdl.Part = input.Part
	_, err = service.repository.UpdateById(c, &lpMdl)
	if err != nil {
		return
	}
	lpRes.FromModel(lpMdl)
	return
}

// Delete
func (service *LecturePartService) Delete(c *fiber.Ctx, input requests.LecturePartDelete) (err error) {
	var lpMdl models.LecturePart
	lpMdl, err = service.repository.FindByModel(c, models.LecturePart{
		Id:       uuid.MustParse(input.Id),
		CourseId: uuid.MustParse(input.CourseId),
	})
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	_, err = service.repository.DeleteByIds(c, &lpMdl)
	if err != nil {
		return
	}
	return
}
