package config

import storage "github.com/ramadani/go-filestorage"

// FileSystemDisks (FileSystemDisks) represents a disk configuration
var FileSystemDisks = map[string]storage.Config{
	"public": {
		Root: "./public",
		URL:  "/public",
	},
}

// GetFileSystemDiskKeys returns a list of available file system disk keys
func GetFileSystemDiskKeys() []string {
	var keys []string
	for key := range FileSystemDisks {
		keys = append(keys, key)
	}
	return keys
}
