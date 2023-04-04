package bootstrap

import (
	"os"
	"regexp"

	"github.com/arfanxn/coursefan-gofiber/config"
)

// FileSystem will bootstraping application filesystem
func FileSystem() error {
	appUrlString := os.Getenv("APP_URL")
	for key, disk := range config.FileSystemDisks {
		if regexp.MustCompile("^/").MatchString(disk.URL) {
			disk.URL = appUrlString + disk.URL
			config.FileSystemDisks[key] = disk
		}
	}
	return nil
}
