package models

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Media struct {
	// Id will be generated automatically, can be set manually if needed
	Id uuid.UUID `json:"id" gorm:"primaryKey;type:char(36)"`
	// ModelType must be specified
	ModelType string `json:"model_type" gorm:"index;type:VARCHAR(25) NOT NULL"`
	// ModelId must be specified
	ModelId uuid.UUID `json:"model_id" gorm:"type:VARCHAR(36) NOT NULL"`
	// CollectionName will be autofilled with default CollectionName if not specified
	CollectionName null.String `json:"collection_name"  gorm:"index;type:VARCHAR(50)"`
	// Name can be null if not specified
	Name null.String `json:"name" gorm:"type:VARCHAR(100)"`
	// FileName will be autofilled by random alphanumeric characters if not specified
	FileName string `json:"file_name" gorm:"type:VARCHAR(256)  NOT NULL"`
	// MimeType will be autofilled by guessing the file bytes if not specified
	MimeType string `json:"mime_type" gorm:"type:VARCHAR(50) NOT NULL"`
	// Disk will be autofilled with default disk if not specified
	Disk string `json:"disk" gorm:"type:VARCHAR(25) NOT NULL"`
	// ConversionDisk can be null if not set
	ConversionDisk null.String `json:"conversion_disk" gorm:"type:VARCHAR(25)"`
	// if not set will be autofilled with by guessing the file bytes size
	Size int64 `json:"size" gorm:"type:INTEGER UNSIGNED NOT NULL"`
	// CreatedAt will be autofilled on creation
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;NOT NULL"`
	// UpdatedAt will be autofilled after updation
	UpdatedAt null.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// FileHeader Metadata, not in table columns
	FileHeader *multipart.FileHeader `json:"-" gorm:"-"`
	// File Metadata, not in table columns
	File *os.File `json:"-" gorm:"-"`
}

// GetFilePath returns media file path based on media disk
func (media *Media) GetFilePath() string {
	fileSystemDisk := config.FileSystemDisks[media.GetDisk()]
	filepath := fmt.Sprintf("%s/medias/%s/%s", fileSystemDisk.Root, media.Id.String(), media.GetFileName())
	return filepath
}

// GetFilePath returns url string to the media file based on media disk
func (media *Media) GetFileUrl() string {
	fileSystemDisk := config.FileSystemDisks[media.GetDisk()]
	urlStr := fmt.Sprintf("%s/%s/medias/%s/%s",
		strings.Trim(os.Getenv("APP_URL"), "/"),
		strings.Trim(fileSystemDisk.URL, "/"),
		media.Id.String(),
		media.GetFileName(),
	)
	return urlStr
}

// GetFileName returns media.FileName , example: image.png
func (media *Media) GetFileName() string {
	if media.FileName != "" {
		return media.FileName
	}
	if media.FileHeader != nil {
		media.FileName = path.Base(media.FileHeader.Filename)
		return media.FileName
	}
	if media.File != nil {
		media.FileName = path.Base(media.File.Name())
		return media.FileName
	}
	return ""
}

// SetFileName sets media.FileName
func (media *Media) SetFileName(fileName string) {
	media.FileName = path.Base(fileName) + path.Ext(media.GetFileName())
}

// GetMimeType returns media.MimeType
func (media *Media) GetMimeType() (string, error) {
	if media.MimeType != "" {
		return media.MimeType, nil
	}
	if media.FileHeader != nil {
		file, err := media.FileHeader.Open()
		if err != nil {
			return "", err
		}
		defer file.Close()
		mime, err := mimetype.DetectReader(file)
		if err != nil {
			return "", err
		}
		media.MimeType = mime.String()
	}
	if media.File != nil {
		mime, err := mimetype.DetectReader(media.File)
		if err != nil {
			return "", err
		}
		media.MimeType = mime.String()
	}
	return media.MimeType, nil
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
func (media *Media) GetSize() (int64, error) {
	if media.Size > 0 {
		return media.Size, nil
	}
	if media.FileHeader != nil {
		media.Size = media.FileHeader.Size
	}
	if media.File != nil {
		fileStat, err := media.File.Stat()
		if err != nil {
			return 0, nil
		}
		media.Size = fileStat.Size()
	}
	return media.Size, nil
}

// SetFileHeader sets media.FileHeader
func (media *Media) SetFileHeader(fh *multipart.FileHeader) error {
	media.FileHeader = fh
	media.FileName = path.Base(fh.Filename)
	media.Size = fh.Size
	mediaMimeType, err := media.GetMimeType()
	if err != nil {
		return err
	}
	media.MimeType = mediaMimeType
	return nil
}

// SetFile sets media.File
func (media *Media) SetFile(file *os.File) (err error) {
	media.File = file
	fileStat, err := file.Stat()
	if err != nil {
		return
	}
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return
	}
	media.FileName = path.Base(file.Name())
	media.Size = fileStat.Size()
	media.MimeType = mime.String()

	// Reset file pointer
	_, err = file.Seek(0, io.SeekStart)
	return nil
}
