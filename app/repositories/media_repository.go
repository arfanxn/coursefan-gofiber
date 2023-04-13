package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/fileh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
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

// Find finds model by id
func (repository *MediaRepository) FindById(c *fiber.Ctx, id string) (media models.Media, err error) {
	err = repository.db.Where("id = ?", id).First(&media).Error
	return
}

// FindByModel finds model by model
func (repository *MediaRepository) FindByModel(c *fiber.Ctx, model models.Media) (media models.Media, err error) {
	err = repository.db.First(&media, model).Error
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
	err := result.Error
	if err != nil {
		fileh.BatchRemove(savedFilePaths...)
		return 0, err
	}
	return result.RowsAffected, nil
}

// UpdateById updates model in the database by given id
func (repository *MediaRepository) UpdateById(c *fiber.Ctx, media *models.Media) (int64, error) {
	// refresh model updated at
	media.UpdatedAt = null.NewTime(time.Now(), true)

	// if the file is also updated
	if media.FileHeader != nil {
		err := fileh.BatchRemove(media.GetFilePath())
		if err != nil {
			return 0, err
		}

		mediaMimeType, err := media.GetMimeType()
		if err != nil {
			return 0, err
		}
		mediaSize, err := media.GetSize()
		if err != nil {
			return 0, err
		}
		media.FileName = media.GetFileName()
		media.Disk = media.GetDisk()
		media.Size = mediaSize
		media.MimeType = mediaMimeType

		mediaFilePath := media.GetFilePath()
		file, err := media.FileHeader.Open()
		if err != nil {
			return 0, err
		}
		defer file.Close()
		err = fileh.Save(mediaFilePath, file)
		if err != nil {
			return 0, err
		}
	}

	// update
	result := repository.db.Model(media).Updates(media)
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
