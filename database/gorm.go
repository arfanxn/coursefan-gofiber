package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
)

// InitGormDB initializes database with configuration
func InitGormDB() (*gorm.DB, error) {
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	config := &gorm.Config{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetGormDB returns the singleton of *gorm.DB or error if it fails
func GetGormDB() (*gorm.DB, error) {
	if gormDB != nil {
		return gormDB, nil
	}
	var err error
	gormDB, err = InitGormDB()
	return gormDB, err
}

// MustGetGormDB returns the singleton of *gorm.DB or panic
func MustGetGormDB() *gorm.DB {
	var err error
	gormDB, err = GetGormDB()
	if err != nil {
		panic(err)
	}
	return gormDB
}
