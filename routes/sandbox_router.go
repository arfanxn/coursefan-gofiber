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
		return
	})
}
