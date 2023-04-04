package seeders

import (
	"bytes"
	"os"
	"path"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/fileh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type MediaSeeder struct {
	repository        *repositories.MediaRepository
	lectureRepository *repositories.LectureRepository
	placeholdersPath  string
}

// NewMediaSeeder instantiates a new MediaSeeder
func NewMediaSeeder(
	repository *repositories.MediaRepository,
	lectureRepository *repositories.LectureRepository,
) *MediaSeeder {
	return &MediaSeeder{
		repository:        repository,
		lectureRepository: lectureRepository,
		placeholdersPath:  config.FileSystemDisks[os.Getenv("MEDIA_DISK")].Root + "/placeholders",
	}
}

// Run runs the seeder
func (seeder *MediaSeeder) Run(c *fiber.Ctx) (err error) {
	// Refresh media directory
	err = fileh.BatchRemove(path.Join(config.FileSystemDisks[os.Getenv("MEDIA_DISK")].Root, "medias"))
	if err != nil {
		return
	}

	errs := []error{
		seeder.SeedLecturesVideos(c),
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return
}

// SeedLecturesVideos seeds video for each models.Lectures
func (seeder *MediaSeeder) SeedLecturesVideos(c *fiber.Ctx) (err error) {
	var medias []*models.Media
	lectures, err := seeder.lectureRepository.All(c)
	if err != nil {
		return
	}
	videoBytes, err := os.ReadFile(path.Join(seeder.placeholdersPath, "video.3gp"))
	if err != nil {
		return err
	}
	for _, lecture := range lectures {
		media := factories.FakeMedia()
		media.Id = uuid.New()
		media.ModelType = reflecth.GetTypeName(lecture)
		media.ModelId = lecture.Id
		media.CollectionName = null.NewString(enums.MediaCollectionNameLectureVideo, true)
		media.SetFileBuffer(bytes.NewBuffer(videoBytes))
		media.SetFileName(faker.Word())
		medias = append(medias, &media)
	}

	for _, chunk := range sliceh.Chunk(medias, 500) {
		_, err = seeder.repository.Insert(c, chunk...)
	}
	return
}
