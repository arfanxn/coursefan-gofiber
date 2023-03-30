package routes

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/arfanxn/coursefan-gofiber/bootstrap"
	"github.com/gofiber/fiber/v2"
)

var (
	testApp              *fiber.App = nil
	testENVConfigured               = false
	testLoggerConfigured            = false
)

// setDirToRoot sets current working directory to  application root directory
func setDirToRoot() error {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	rootdir := filepath.Dir(d)
	return os.Chdir(rootdir)
}

// setup sets up the test
func setupTest() (err error) {
	setDirToRoot()

	if testENVConfigured == false {
		err = bootstrap.ConfigureTestENV()
		if err != nil {
			return
		}
		testENVConfigured = true
	}

	if testLoggerConfigured == false {
		err = bootstrap.ConfigureLogger()
		if err != nil {
			return err
		}
		testLoggerConfigured = true
	}

	if testApp == nil {
		testApp, err = bootstrap.NewApp()
		if err != nil {
			testApp = nil
			return err
		}
		InjectRoutes(testApp)
	}

	return nil
}
