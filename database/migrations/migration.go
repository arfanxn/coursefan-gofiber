package migrations

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/database"
)

func MigrateUp() error {
	db, err := database.GetGORM()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Media{},
	)
	if err != nil {
		return err
	}

	return nil
}
