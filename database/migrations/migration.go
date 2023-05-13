package migrations

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
)

var tables []any = []any{
	&models.Media{},
	&models.Token{},
	&models.User{},
	&models.UserProfile{},
	&models.UserSetting{},
	&models.Notification{},
	&models.Message{},
	&models.Permission{},
	&models.Role{},
	&models.PermissionRole{},
	&models.Course{},
	&models.CourseOrder{},
	&models.CourseUserRole{},
	&models.Review{},
	&models.LecturePart{},
	&models.Lecture{},
	&models.Discussion{},
	&models.Progress{},
	&models.ProgressUser{},
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
