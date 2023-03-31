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
	testApp *fiber.App = nil
)

func TestMain(m *testing.M) {
	// Set testing directory to root application directory
	setDirToRoot()

	// Load test environment variables
	err := godotenv.Load(config.TestEnvironmentFileName)
	if err != nil {
		logrus.Fatal(err)
	}
	// Boot Logger
	err = bootstrap.Logger()
	if err != nil {
		logrus.Fatal(err)
	}

	// Boot new Application and assign it to testApp variable
	testApp, err = bootstrap.NewApp()
	if err != nil {
		logrus.Fatal(err)
	}

	RegisterApp(testApp)

	// Migrate up required tables
	err = migrations.MigrateUp()
	if err != nil {
		logrus.Fatal(err)
	}

	// Run tests
	exitCode := m.Run()

	// Migrate down
	// migrations.MigrateDown()
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	os.Exit(exitCode)
}

// setDirToRoot sets current working directory to  application root directory
func setDirToRoot() error {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	rootdir := filepath.Dir(d)
	return os.Chdir(rootdir)
}
