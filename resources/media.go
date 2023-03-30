package resources

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"gopkg.in/guregu/null.v4"
)

type Media struct {
	Id             string      `json:"id"`
	ModelType      string      `json:"model_type"`
	ModelId        string      `json:"model_id"`
	CollectionName null.String `json:"collection_name"`
	Name           null.String `json:"name"`
	FileName       string      `json:"file_name"`
	FileUrl        string      `json:"file_url"` // url to the file
	MimeType       string      `json:"mime_type"`
	Disk           string      `json:"disk"`
	ConversionDisk null.String `json:"conversion_disk"`
	Size           int64       `json:"size"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      null.Time   `json:"updated_at"`
}

// NewMediaFromModel instantiates resources.Media with values from the given models.Media
func NewMediaFromModel(mediaMdl models.Media) *Media {
	mediaRes := new(Media)
	mediaRes.Id = mediaMdl.Id.String()
	mediaRes.ModelType = mediaMdl.ModelType
	mediaRes.ModelId = mediaMdl.ModelId.String()
	mediaRes.CollectionName = null.NewString(mediaMdl.CollectionName.String, mediaMdl.CollectionName.Valid)
	mediaRes.Name = null.NewString(mediaMdl.Name.String, mediaMdl.Name.Valid)
	mediaRes.FileName = mediaMdl.GetFileName()
	mediaRes.FileUrl = mediaMdl.GetFileUrl()
	mediaRes.MimeType = mediaMdl.GetMimeType()
	mediaRes.Disk = mediaMdl.GetDisk()
	mediaRes.ConversionDisk = null.NewString(mediaMdl.ConversionDisk.String, mediaMdl.ConversionDisk.Valid)
	mediaRes.Size = mediaMdl.Size
	mediaRes.CreatedAt = mediaMdl.CreatedAt
	mediaRes.UpdatedAt = null.NewTime(mediaMdl.UpdatedAt.Time, mediaMdl.UpdatedAt.Valid)
	return mediaRes
}
