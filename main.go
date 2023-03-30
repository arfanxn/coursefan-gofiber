package main

import (
	"net/url"
	"os"

	"github.com/arfanxn/coursefan-gofiber/app/console"
	"github.com/arfanxn/coursefan-gofiber/bootstrap"
	"github.com/arfanxn/coursefan-gofiber/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	console.Execute() // CMD Execution

	err := bootstrap.ConfigureLogger()
	if err != nil {
		logrus.Fatal(err)
	}

	app, err := bootstrap.NewAppWithConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	err = routes.InjectRoutes(app)
	if err != nil {
		logrus.Fatal(err)
	}

	appURL, err := url.Parse(os.Getenv("APP_URL"))
	if err != nil {
		logrus.Fatal(err)
	}

	err = app.Listen(appURL.Host)
	if err != nil {
		logrus.Fatal(err)
	}
}
