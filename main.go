package main

import (
	"net/url"
	"os"

	"github.com/arfanxn/coursefan-gofiber/app/console"
	"github.com/arfanxn/coursefan-gofiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	console.Execute() // CMD Exception

	// err := logger.SetOutputBasedOnENV()
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	app := fiber.New()

	err := routes.InjectRoutes(app)
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
