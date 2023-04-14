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

type LectureService struct {
	repository      *repositories.LectureRepository
	mediaRepository *repositories.MediaRepository
}

func NewLectureService(
	repository *repositories.LectureRepository,
	mediaRepository *repositories.MediaRepository,
) *LectureService {
	return &LectureService{
		repository:      repository,
		mediaRepository: mediaRepository,
	}
}

// AllByLecturePart get lectures by lecture part
func (service *LectureService) AllByLecturePart(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.Lecture], err error) {
	lectureMdls, err := service.repository.AllByLecturePart(c, input)
	if err != nil {
		return
	}
	lectureRess := sliceh.Map(lectureMdls, func(lectureMdl models.Lecture) resources.Lecture {
		lectureRes := resources.Lecture{}
		lectureRes.FromModel(lectureMdl)
		return lectureRes
	})
	pagination.SetItems(lectureRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), int(input.Limit.Int64))
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *LectureService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.Lecture, err error) {
	lectureMdls, err := service.repository.AllByLecturePart(c, input)
	if err != nil {
		return
	} else if len(lectureMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	lectureMdl := lectureMdls[0]
	lectureVideoMediaMdl, err := service.mediaRepository.FindByModel(c, models.Media{
		ModelId:        lectureMdl.Id,
		ModelType:      reflecth.GetTypeName(lectureMdl),
		CollectionName: null.NewString(enums.MediaCollectionNameLectureVideo, true),
	})
	if err != nil {
		return
	}

	data.FromModel(lectureMdl)
	lectureVideoMediaRes := resources.Media{}
	lectureVideoMediaRes.FromModel(lectureVideoMediaMdl)
	data.Video = &lectureVideoMediaRes

	return
}

// Create
func (service *LectureService) Create(c *fiber.Ctx, input requests.LectureCreate) (
	lectureRes resources.Lecture, err error) {
	lectureMdl := models.Lecture{}
	lectureMdl.LecturePartId = uuid.MustParse(input.LecturePartId)
	lectureMdl.Name = input.Name
	lectureMdl.Order = input.Order
	_, err = service.repository.Insert(c, &lectureMdl)
	if err != nil {
		return
	}

	// Save the lecture video
	lectureVideoMediaMdl := models.Media{}
	lectureVideoMediaMdl.ModelId = lectureMdl.Id
	lectureVideoMediaMdl.ModelType = reflecth.GetTypeName(lectureMdl)
	lectureVideoMediaMdl.CollectionName = null.NewString(enums.MediaCollectionNameLectureVideo, true)
	err = lectureVideoMediaMdl.SetFileHeader(input.Video)
	if err != nil {
		return
	}
	_, err = service.mediaRepository.Insert(c, &lectureVideoMediaMdl)
	if err != nil {
		return
	}

	lectureRes.FromModel(lectureMdl)
	lectureVideoMediaRes := resources.Media{}
	lectureVideoMediaRes.FromModel(lectureVideoMediaMdl)
	lectureRes.Video = &lectureVideoMediaRes
	return
}

// Update
func (service *LectureService) Update(c *fiber.Ctx, input requests.LectureUpdate) (
	lectureRes resources.Lecture, err error) {

	lectureMdl, err := service.repository.FindById(c, input.Id)
	if errorh.IsGormErrRecordNotFound(err) {
		err = fiber.ErrNotFound
		return
	} else if err != nil {
		return
	}
	lectureMdl.Name = input.Name
	lectureMdl.Order = input.Order
	_, err = service.repository.UpdateById(c, &lectureMdl)
	if err != nil {
		return
	}

	lectureVideoMediaMdl, err := service.mediaRepository.FindByModel(c, models.Media{
		ModelId:        lectureMdl.Id,
		ModelType:      reflecth.GetTypeName(lectureMdl),
		CollectionName: null.NewString(enums.MediaCollectionNameLectureVideo, true),
	})
	if err != nil {
		return
	}

	if input.Video != nil {
		err = lectureVideoMediaMdl.SetFileHeader(input.Video)
		if err != nil {
			return
		}
		_, err = service.mediaRepository.UpdateById(c, &lectureVideoMediaMdl)
		if err != nil {
			return
		}
	}

	lectureRes.FromModel(lectureMdl)
	lectureVideoMediaRes := resources.Media{}
	lectureVideoMediaRes.FromModel(lectureVideoMediaMdl)
	lectureRes.Video = &lectureVideoMediaRes
	return
}

// Delete
func (service *LectureService) Delete(c *fiber.Ctx, input requests.LectureDelete) (err error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	syncronizer.WG().Add(2)

	go func() {
		defer syncronizer.WG().Done()
		var lectureMdl models.Lecture
		lectureMdl, err = service.repository.FindById(c, input.Id)
		if errorh.IsGormErrRecordNotFound(err) {
			err = fiber.ErrNotFound
			syncronizer.Err(err)
			return
		} else if err != nil {
			syncronizer.Err(err)
			return
		}
		_, err = service.repository.DeleteByIds(c, &lectureMdl)
		if err != nil {
			syncronizer.Err(err)
			return
		}
	}()
	go func() {
		defer syncronizer.WG().Done()
		lectureVideoMediaMdl, err := service.mediaRepository.FindByModel(c, models.Media{
			ModelId:        uuid.MustParse(input.Id),
			ModelType:      reflecth.GetTypeName(models.Lecture{}),
			CollectionName: null.NewString(enums.MediaCollectionNameLectureVideo, true),
		})
		if err != nil {
			syncronizer.Err(err)
			return
		}
		_, err = service.mediaRepository.DeleteByIds(c, &lectureVideoMediaMdl)
		if err != nil {
			syncronizer.Err(err)
			return
		}
	}()
	syncronizer.WG().Wait()
	if err = syncronizer.Err(); err != nil {
		return
	}
	return
}
