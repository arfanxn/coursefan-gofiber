package services

import (
	"fmt"
	"net/url"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/numh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
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

// Create
func (service *CourseService) Create(c *fiber.Ctx, input requests.CourseCreate) (
	courseRes resources.Course, err error) {
	courseMdl := models.Course{}
	courseMdl.Name = input.Name
	courseMdl.Description = input.Description
	courseMdl.Slug = fmt.Sprintf("%s-%d", strcase.ToKebab(courseMdl.Name), numh.Random(1000, 9999))
	_, err = service.repository.Insert(c, &courseMdl)
	if err != nil {
		return
	}
	courseRes = resources.Course{}
	courseRes.FromModel(courseMdl)
	return
}
