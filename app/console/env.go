package console

import (
	"github.com/arfanxn/coursefan-gofiber/config"
)

var (
	envFilename *string
)

func envFlag() (err error) {
	// Change the environment variable configuration
	config.EnvironmentFileName = *envFilename
	if err != nil {
		return
	}
	return
}

func init() {
	envFilename = rootCmd.Flags().StringP("env", "e", config.EnvironmentFileName, "The environment variable")
}
