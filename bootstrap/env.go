package bootstrap

import (
	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/joho/godotenv"
)

// ENV will bootstraping application environment variables
func ENV() (err error) {
	err = godotenv.Load(config.EnvironmentFileName)
	return
}
