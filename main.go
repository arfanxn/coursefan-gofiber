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

	err := bootstrap.Logger()
	if err != nil {
		logrus.Fatal(err)
	}

	err = bootstrap.Validator()
	if err != nil {
		logrus.Fatal(err)
	}

	app, err := bootstrap.NewApp()
	if err != nil {
		logrus.Fatal(err)
	}

	routes.InjectRoutes(app)

	appURL, err := url.Parse(os.Getenv("APP_URL"))
	if err != nil {
		logrus.Fatal(err)
	}

	err = app.Listen(appURL.Host)
	if err != nil {
		logrus.Fatal(err)
	}
}
