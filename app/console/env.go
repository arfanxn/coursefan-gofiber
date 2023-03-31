package console

import (
	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/joho/godotenv"
)

var (
	envFilename *string
)

func envFlag() (err error) {
	// Load env variables based on specified env filename
	err = godotenv.Load(*envFilename)
	if err != nil {
		return
	}
	return
}

func init() {
	envFilename = rootCmd.Flags().StringP("env", "e", config.EnvironmentFileName, "The environment variable")
}
