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
