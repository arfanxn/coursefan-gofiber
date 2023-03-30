package bootstrap

import "github.com/joho/godotenv"

// ConfigureTestENV will bootstrap environment variables with test environment into application
func ConfigureTestENV() error {
	testEnvFileName := "test.env"
	return godotenv.Load(testEnvFileName)
}
