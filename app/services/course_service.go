package services

import (
	"fmt"

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
	return
}

// Find
func (service *CourseService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.Course, err error) {
	courseMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	} else if len(courseMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	data.FromModel(courseMdls[0])
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

// Update
func (service *CourseService) Update(c *fiber.Ctx, input requests.CourseUpdate) (
	courseRes resources.Course, err error) {
	courseMdl, err := service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	courseMdl.Name = input.Name
	courseMdl.Description = input.Description
	_, err = service.repository.UpdateById(c, &courseMdl)
	if err != nil {
		return
	}
	courseRes = resources.Course{}
	courseRes.FromModel(courseMdl)
	return
}

// Delete
func (service *CourseService) Delete(c *fiber.Ctx, input requests.CourseDelete) (err error) {
	var courseMdl models.Course
	courseMdl, err = service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	_, err = service.repository.DeleteByIds(c, &courseMdl)
	if err != nil {
		return
	}
	return
}
