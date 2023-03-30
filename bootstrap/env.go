package bootstrap

import "github.com/joho/godotenv"

// UseTestEnv will bootstrap environment variables with test environment
func UseTestEnv() error {
	testEnvFileName := "test.env"
	return godotenv.Load(testEnvFileName)
}
