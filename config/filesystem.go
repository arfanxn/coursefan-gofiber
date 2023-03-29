package config

type fileSystemDisk struct {
	Root string // the root of the filesystem
	URL  string // the URL to access file
}

// FileSystemDisks (FileSystemDisks) represents a disk configuration
var FileSystemDisks = map[string]fileSystemDisk{
	"public": {
		Root: "./public/",
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
