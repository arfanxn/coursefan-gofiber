package bootstrap

import (
	"os"

	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/joho/godotenv"
)

// ENV will bootstraping application environment variables
func ENV() (err error) {
	// Check if environment is loaded before bootstrap, if already loaded then immediately return
	if os.Getenv("APP_NAME") != "" {
		return
	}
	err = godotenv.Load(config.EnvironmentFileName)
	return
}
