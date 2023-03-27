package migrations

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	conn "github.com/arfanxn/coursefan-gofiber/database/connection"
)

func MigrateUp() error {
	db, err := conn.GetGORM()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}
