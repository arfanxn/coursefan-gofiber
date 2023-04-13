// databasep stands for database provider
package databasep

import (
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/bootstrap"
	"gorm.io/gorm"
)

var (
	// gormDB instance
	gormDB *gorm.DB = nil
)

// GetGormDB returns the singleton of *gorm.DB or error if it fails
func GetGormDB() (*gorm.DB, error) {
	syncronizer := synch.NewSyncronizer()
	defer syncronizer.Close()
	syncronizer.RWM().Lock()
	defer syncronizer.RWM().Unlock()
	if gormDB != nil {
		return gormDB, nil
	}
	var err error
	gormDB, err = bootstrap.NewGormDB()
	return gormDB, err
}

// MustGetGormDB returns the singleton of *gorm.DB or panic
func MustGetGormDB() *gorm.DB {
	var (
		gormDB *gorm.DB
		err    error
	)
	gormDB, err = GetGormDB()
	if err != nil {
		panic(err)
	}
	return gormDB
}
