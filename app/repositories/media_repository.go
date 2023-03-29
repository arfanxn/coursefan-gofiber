package repositories

import (
	"time"

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
func (repository *MediaRepository) Insert(c *fiber.Ctx, medias ...*models.Media) (err error) {
	syncronizer := synch.NewSyncronizer()
	for _, media := range medias {
		syncronizer.WG().Add(1)
		go func(media *models.Media) {
			defer syncronizer.WG().Done()
			if media.Id == uuid.Nil {
				media.Id = uuid.New()
			}

			media.FileName = media.GetFileName()
			media.Disk = media.GetDisk()
			media.Size = media.GetSize()
			media.MimeType = media.GetMimeType()

			media.CreatedAt = time.Now()
		}(media)
	}
	syncronizer.WG().Wait()

	err = repository.db.Create(medias).Error
	return
}
