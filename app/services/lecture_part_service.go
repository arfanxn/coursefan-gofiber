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

// GetByCourse get lecture parts by course
func (service *LecturePartService) GetByCourse(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.LecturePart], err error) {
	courseMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	}
	lecturePartRess := sliceh.Map(courseMdls, func(lpMdl models.LecturePart) resources.LecturePart {
		lpRes := resources.LecturePart{}
		lpRes.FromModel(lpMdl)
		return lpRes
	})
	pagination.SetItems(lecturePartRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), input.Limit)
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}
