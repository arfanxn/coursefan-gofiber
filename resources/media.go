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

// FromModel instantiates resources.Media with values from the given models.Media
func (resource *Media) FromModel(model models.Media) error {
	mediaMimeType, err := model.GetMimeType()
	if err != nil {
		return err
	}

	resource.Id = model.Id.String()
	resource.ModelType = model.ModelType
	resource.ModelId = model.ModelId.String()
	resource.CollectionName = null.NewString(model.CollectionName.String, model.CollectionName.Valid)
	resource.Name = null.NewString(model.Name.String, model.Name.Valid)
	resource.FileName = model.GetFileName()
	resource.FileUrl = model.GetFileUrl()
	resource.MimeType = mediaMimeType
	resource.Disk = model.GetDisk()
	resource.ConversionDisk = null.NewString(model.ConversionDisk.String, model.ConversionDisk.Valid)
	resource.Size = model.Size
	resource.CreatedAt = model.CreatedAt
	resource.UpdatedAt = null.NewTime(model.UpdatedAt.Time, model.UpdatedAt.Valid)

	return nil
}
