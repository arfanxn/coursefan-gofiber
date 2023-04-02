//go:build wireinject
// +build wireinject

// seederp  stands for Seeder Provider
package seederp

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/seeders"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitUserSeeder(db *gorm.DB) *seeders.UserSeeder {
	wire.Build(
		repositories.NewUserRepository,
		seeders.NewUserSeeder,
	)
	return nil
}
