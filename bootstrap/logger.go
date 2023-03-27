package bootstrap

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// SetLoggerOutputBasedOnENV sets logger output globally based on the environment variable LOG_OUTPUT
func SetLoggerOutputBasedOnENV() error {
	// Get log file output name from environment variable
	logFilename := os.Getenv("LOG_OUTPUT")
	logFilename = "storage" + "/" + logFilename

	// Create log directory if it doesn't exist
	err := os.MkdirAll(path.Dir(logFilename), os.ModePerm)
	if err != nil {
		return err
	}

	// Open log output file
	logFile, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	// Set log output file as logger output globally
	logrus.SetOutput(logFile)

	return nil
}
