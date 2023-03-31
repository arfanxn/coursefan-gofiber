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

	// Boot Logger
	err := bootstrap.Logger()
	if err != nil {
		logrus.Fatal(err)
	}
	// Boot new Application
	app, err := bootstrap.NewApp()
	if err != nil {
		logrus.Fatal(err)
	}
	// inject/register routes into application instance
	routes.RegisterApp(app)

	appURL, err := url.Parse(os.Getenv("APP_URL"))
	if err != nil {
		logrus.Fatal(err)
	}

	// Start listening server
	err = app.Listen(appURL.Host)
	if err != nil {
		logrus.Fatal(err)
	}
}
