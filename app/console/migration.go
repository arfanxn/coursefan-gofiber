package console

import "github.com/arfanxn/coursefan-gofiber/database/migrations"

var (
	migrateUp   *bool
	migrateDown *bool
)

func migrateFlag() (err error) {
	if *migrateUp == true {
		err = migrations.MigrateUp()
	}

	// TODO: implement migrate down

	return
}

func init() {
	migrateUp = rootCmd.Flags().Bool("migrate-up", false, "Migrate up database tables")
	migrateDown = rootCmd.Flags().Bool("migrate-down", false, "Migrate down database tables")
}
