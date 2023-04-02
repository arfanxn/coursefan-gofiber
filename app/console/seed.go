package console

import (
	"fmt"

	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/arfanxn/coursefan-gofiber/database/seeders"
	"github.com/gofiber/fiber/v2"
)

var (
	seed *bool
)

func seedFlag() (err error) {
	if *seed == true {
		c := new(fiber.Ctx)
		dbs := seeders.NewDatabaseSeeder(databasep.MustGetGormDB())
		err = dbs.Run(c)
		exitAfterFinish = true
		if err != nil {
			return
		}
		fmt.Println("Successfully seed database")
	}

	return
}

func init() {
	seed = rootCmd.Flags().Bool("seed", false, "Seed the database")
}
