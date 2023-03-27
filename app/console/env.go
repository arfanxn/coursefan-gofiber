package console

import (
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
	defaultEnvFilename := "local.env"
	envFilename = rootCmd.Flags().StringP("env", "e", defaultEnvFilename, "The environment variable")
}
