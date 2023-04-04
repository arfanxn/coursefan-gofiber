package routes

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/arfanxn/coursefan-gofiber/bootstrap"
	"github.com/arfanxn/coursefan-gofiber/config"
	"github.com/arfanxn/coursefan-gofiber/database/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	testApp                *fiber.App = nil
	testErr                error      = nil
	testHTTPRequestTimeout int        = -1
)

func TestMain(m *testing.M) {
	// Set testing directory to root application directory
	testErr = setDirToRoot()
	if testErr != nil {
		logrus.Fatal(testErr)
	}

	// Load test environment variables
	testErr = godotenv.Load(config.TestEnvironmentFileName)
	if testErr != nil {
		logrus.Fatal(testErr)
	}
	// Boot everything
	bootstrap.Boot()
	// Boot new Application and assign it to testApp variable
	testApp, testErr = bootstrap.NewApp()
	if testErr != nil {
		logrus.Fatal(testErr)
	}
	// inject/register routes into application instance
	RegisterApp(testApp)

	// Migrate up required tables
	testErr = migrations.MigrateUp()
	if testErr != nil {
		logrus.Fatal(testErr)
	}

	// Run tests
	exitCode := m.Run()

	// Migrate down after tests is finished
	migrations.MigrateDown()
	if testErr != nil {
		logrus.Fatal(testErr)
	}

	os.Exit(exitCode)
}

// setDirToRoot sets current working directory to  application root directory
func setDirToRoot() error {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	rootdir := filepath.Dir(d)
	return os.Chdir(rootdir)
}
