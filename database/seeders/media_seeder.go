package seeders

import (
	"os"
	"path"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/fileh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
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
	// TODO: fix save file courrupted

	var medias []*models.Media
	lectures, err := seeder.lectureRepository.All(c)
	if err != nil {
		return
	}
	videoFile, err := os.Open(path.Join(seeder.placeholdersPath, "potrait.jpg"))
	if err != nil {
		return err
	}
	for _, lecture := range lectures {
		media := factories.FakeMedia()
		media.ModelType = reflecth.GetTypeName(lecture)
		media.ModelId = lecture.Id
		err = media.SetFile(videoFile)
		if err != nil {
			return err
		}
		medias = append(medias, &media)
	}
	_, err = seeder.repository.Insert(c, medias...)
	return
}
