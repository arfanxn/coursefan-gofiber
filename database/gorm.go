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

// InitGORM initializes database with configuration
func InitGORM() (*gorm.DB, error) {
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	config := &gorm.Config{}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetGORM returns the singleton of *gorm.DB or error if it fails
func GetGORM() (*gorm.DB, error) {
	if gormDB != nil {
		return gormDB, nil
	}
	var err error
	gormDB, err = InitGORM()
	return gormDB, err
}

// MustGetGORM returns the singleton of *gorm.DB or panic
func MustGetGORM() *gorm.DB {
	if gormDB != nil {
		return gormDB
	}
	var err error
	gormDB, err = InitGORM()
	if err != nil {
		panic(err)
	}
	return gormDB
}
