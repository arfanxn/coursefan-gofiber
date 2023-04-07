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

type CourseService struct {
	repository *repositories.CourseRepository
}

// NewCourseService instantiates a new CourseService
func NewCourseService(
	repository *repositories.CourseRepository,
) *CourseService {
	return &CourseService{
		repository: repository,
	}
}

// All
func (service *CourseService) All(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.Course], err error) {
	courseMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	}
	courseRess := sliceh.Map(courseMdls, func(courseMdl models.Course) resources.Course {
		courseRes := resources.Course{}
		courseRes.FromModel(courseMdl)
		return courseRes
	})
	pagination.SetItems(courseRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), input.Limit)
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))

	return
}
