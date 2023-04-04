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

// All returns all medias in the database
func (repository *MediaRepository) All(c *fiber.Ctx) (medias []models.Media, err error) {
	err = repository.db.Find(&medias).Error
	return
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

			mediaMimeType, err := media.GetMimeType()
			if err != nil {
				syncronizer.Err(err)
				return
			}
			mediaSize, err := media.GetSize()
			if err != nil {
				syncronizer.Err(err)
				return
			}

			media.FileName = media.GetFileName()
			media.Disk = media.GetDisk()
			media.Size = mediaSize
			media.MimeType = mediaMimeType

			media.CreatedAt = time.Now()

			mediaFilePath := media.GetFilePath()
			if media.FileHeader != nil {
				file, err := media.FileHeader.Open()
				if err != nil {
					syncronizer.Err(err)
					return
				}
				defer file.Close()
				err = fileh.Save(mediaFilePath, file)
				if err != nil {
					syncronizer.Err(err)
					return
				}
			}
			if media.FileBuffer != nil {
				err = fileh.Save(mediaFilePath, media.FileBuffer)
				if err != nil {
					syncronizer.Err(err)
					return
				}
			}

			syncronizer.M().Lock()
			defer syncronizer.M().Unlock()
			savedFilePaths = append(savedFilePaths, mediaFilePath)
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

// DeleteByIds deletes the entities associated with the given ids
func (repository *MediaRepository) DeleteByIds(c *fiber.Ctx, medias ...*models.Media) (int64, error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()

	for _, media := range medias {
		syncronizer.WG().Add(1)
		go func(media *models.Media) {
			defer syncronizer.WG().Done()
			if syncronizer.Err() != nil {
				return
			}
			syncronizer.Err(
				fileh.BatchRemove(media.GetFilePath()),
			)
		}(media)
	}
	syncronizer.WG().Wait()
	if err := syncronizer.Err(); err != nil {
		return 0, err
	}

	result := repository.db.Delete(medias)
	return result.RowsAffected, result.Error
}
