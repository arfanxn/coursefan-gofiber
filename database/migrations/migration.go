package migrations

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
)

var tables []any = []any{
	&models.User{},
	&models.UserSetting{},
	&models.Media{},
	&models.Token{},
}

func MigrateUp() error {
	db, err := databasep.GetGormDB()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(tables...)
	if err != nil {
		return err
	}

	return nil
}

func MigrateDown() error {
	db, err := databasep.GetGormDB()
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(tables...)
	if err != nil {
		return err
	}

	return nil
}
