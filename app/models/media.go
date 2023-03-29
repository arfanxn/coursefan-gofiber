package models

import (
	"database/sql"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Media struct {
	// Id will be generated automatically, can be set manually if needed
	Id uuid.UUID `json:"id"`
	// ModelType must be specified
	ModelType string `json:"model_type"`
	// ModelId must be specified
	ModelId uuid.UUID `json:"model_id"`
	// CollectionName will be autofilled with default CollectionName if not specified
	CollectionName string `json:"collection_name"`
	// Name can be null if not specified
	Name sql.NullString `json:"name"`
	// FileName will be autofilled by random alphanumeric characters if not specified
	FileName string `json:"file_name"`
	// MimeType will be autofilled by guessing the file bytes if not specified
	MimeType string `json:"mime_type"`
	// Disk will be autofilled with default disk if not specified
	Disk string `json:"disk"`
	// ConversionDisk can be null if not set
	ConversionDisk sql.NullString `json:"conversion_disk"`
	// if not set will be autofilled with by guessing the file bytes size
	Size int64 `json:"size"`
	// CreatedAt will be autofilled on creation
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt will be autofilled after updation
	UpdatedAt sql.NullTime `json:"updated_at"`

	// File Metadata, not in table columns
	fileHeader *multipart.FileHeader `json:"-"`

	// Model Relation
	Model any `json:"model"`
}
