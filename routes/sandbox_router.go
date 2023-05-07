package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/davecgh/go-spew/spew"
	"github.com/gofiber/fiber/v2"
)

func registerSandboxRouter(router fiber.Router) {
	router.Get("/sandbox", func(c *fiber.Ctx) (err error) {
		input := requests.Query{}
		err = input.FromContext(c)
		// spew.Dump(input)
		var userMdls []models.User
		db := databasep.MustGetGormDB()
		db = gormh.BuildFromRequestQuery(db, models.User{}, input)
		db = db.Find(&userMdls)
		err = db.Error
		if err != nil {
			return
		}
		spew.Dump("ERROR: ", err)
		spew.Dump(userMdls)

		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		err = c.SendString(`<html>
		<body>
		<video width="320" height="240" controls>
		<source src="http://localhost:8080/public/medias/dc133137-8287-432a-9aef-cd53ae71fbb6/sample_3840x2160.mp4" type="video/mp4">
	  Your browser does not support the video tag.
	  </video>
		</body>
		</html>
		`)
		if err != nil {
			return
		}

		return
	})
}
