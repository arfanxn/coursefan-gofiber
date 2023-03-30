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
	if *migrateUp == true {
		err = migrations.MigrateUp()
		fmt.Println("Successfully migrate up database")
		exitAfterFinish = true
	}

	if *migrateDown == true {
		err = migrations.MigrateDown()
		fmt.Println("Successfully migrate down database")
		exitAfterFinish = true
	}

	// TODO: implement migrate down

	return
}

func init() {
	migrateUp = rootCmd.Flags().Bool("migrate-up", false, "Migrate up database tables")
	migrateDown = rootCmd.Flags().Bool("migrate-down", false, "Migrate down database tables")
}
