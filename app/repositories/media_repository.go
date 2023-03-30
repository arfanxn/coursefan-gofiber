package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/fileh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaRepository struct {
	db *gorm.DB
}

// NewMediaRepository instantiates a new MediaRepository
func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{db: db}
}

// Insert inserts medias into the database
func (repository *MediaRepository) Insert(c *fiber.Ctx, medias ...*models.Media) (int64, error) {
	var savedFilePaths []string
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	for _, media := range medias {
		syncronizer.WG().Add(1)
		go func(media *models.Media) {
			defer syncronizer.WG().Done()
			if syncronizer.Err() != nil {
				return
			}

			if media.Id == uuid.Nil {
				media.Id = uuid.New()
			}

			media.FileName = media.GetFileName()
			media.Disk = media.GetDisk()
			media.Size = media.GetSize()
			media.MimeType = media.GetMimeType()

			media.CreatedAt = time.Now()

			file, err := media.FileHeader.Open()
			if err != nil {
				syncronizer.Err(err)
				return
			}
			defer file.Close()

			filePath := media.GetFilePath()
			err = fileh.Save(filePath, file)
			if err != nil {
				syncronizer.Err(err)
				return
			}
			syncronizer.M().Lock()
			defer syncronizer.M().Unlock()
			savedFilePaths = append(savedFilePaths, filePath)
		}(media)
	}
	syncronizer.WG().Wait()

	if err := syncronizer.Err(); err != nil {
		fileh.BatchRemove(savedFilePaths...)
		return 0, err
	}
	result := repository.db.Create(medias)
	return result.RowsAffected, result.Error
}
