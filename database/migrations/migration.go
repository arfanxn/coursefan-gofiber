package migrations

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
)

var tables []any = []any{
	&models.Discussion{},
	&models.Notification{},
	&models.Transaction{},
	&models.Wallet{},
	&models.Media{},
	&models.Review{},
	&models.Token{},
	&models.User{},
	&models.UserSetting{},
	&models.Permission{},
	&models.Role{},
	&models.PermissionRole{},
	&models.Course{},
	&models.CourseUserRole{},
	&models.Lecture{},
	&models.Part{},
}

func MigrateUp() error {
	db, err := databasep.GetGormDB()
	if err != nil {
		return err
	}

	err = db.Migrator().AutoMigrate(tables...)
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
