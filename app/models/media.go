package models

import (
	"database/sql"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/gabriel-vasile/mimetype"
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
	FileHeader *multipart.FileHeader `json:"-"`

	// Model Relation
	Model any `json:"model"`
}

// GetFileName returns media.FileName
func (media *Media) GetFileName() string {
	if media.FileName != "" {
		return media.FileName
	}
	if media.FileHeader != nil {
		media.FileName = path.Base(media.FileHeader.Filename)
		return media.FileName
	}
	return ""
}

// SetFileName sets media.FileName
func (media *Media) SetFileName(fileName string) {
	media.FileName = path.Base(fileName) + path.Ext(media.GetFileName())
}

// GetMimeType returns media.MimeType
func (media *Media) GetMimeType() string {
	if media.MimeType != "" {
		return media.MimeType
	}
	if media.FileHeader != nil {
		file, err := media.FileHeader.Open()
		if err != nil {
			panic(err)
		}
		defer file.Close()
		mime, err := mimetype.DetectReader(file)
		if err != nil {
			panic(err)
		}
		media.MimeType = mime.String()
	}
	return media.MimeType
}

// GetDisk returns media.Disk
func (media *Media) GetDisk() string {
	if media.Disk != "" {
		return media.Disk
	}
	media.Disk = os.Getenv("MEDIA_DISK")
	return media.Disk
}

// SetDisk sets media.Disk
func (media *Media) SetDisk(disk string) {
	disks := config.GetFileSystemDiskKeys()
	// if disk is not empty and disk is available in list of disks
	if disk != "" && sliceh.Contains(disks, disk) {
		media.Disk = disk
		return
	}
	// otherwise assign default media disk to media.Disk
	media.Disk = os.Getenv("MEDIA_DISK")
}

// GetSize returns media.Size
func (media *Media) GetSize() int64 {
	if media.Size > 0 {
		return media.Size
	}
	if media.FileHeader != nil {
		media.Size = media.FileHeader.Size
	}
	return media.Size
}

// SetFileHeader sets media.FileHeader
func (media *Media) SetFileHeader(fh *multipart.FileHeader) {
	media.FileHeader = fh
	media.FileName = path.Base(fh.Filename)
	media.Size = fh.Size
	media.MimeType = media.GetMimeType()
}
