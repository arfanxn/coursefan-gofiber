package console

import (
	"fmt"

	"github.com/arfanxn/coursefan-gofiber/database/migrations"
)

var (
	migrateUp   *bool
	migrateDown *bool
)

func migrateFlag() (err error) {
	if *migrateDown == true {
		err = migrations.MigrateDown()
		exitAfterFinish = true
		if err != nil {
			return
		}
		fmt.Println("Successfully migrate down database")
	}

	if *migrateUp == true {
		err = migrations.MigrateUp()
		exitAfterFinish = true
		if err != nil {
			return
		}
		fmt.Println("Successfully migrate up database")
	}

	return
}

func init() {
	migrateUp = rootCmd.Flags().Bool("migrate-up", false, "Migrate up database tables")
	migrateDown = rootCmd.Flags().Bool("migrate-down", false, "Migrate down database tables")
}
