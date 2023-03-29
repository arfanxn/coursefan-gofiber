package resources

import (
	"database/sql"
	"time"

	"gopkg.in/guregu/null.v4"
)

type Media struct {
	Id             string         `json:"id"`
	ModelType      string         `json:"model_type"`
	ModelId        string         `json:"model_id"`
	CollectionName string         `json:"collection_name"`
	Name           null.String    `json:"name"`
	FileName       string         `json:"file_name"`
	FileUrl        string         `json:"file_url"` // url to the file
	MimeType       string         `json:"mime_type"`
	Disk           string         `json:"disk"`
	ConversionDisk sql.NullString `json:"conversion_disk"`
	Size           int64          `json:"size"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      null.Time      `json:"updated_at"`
}

