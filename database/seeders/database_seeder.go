package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DatabaseSeeder struct {
	db      *gorm.DB
	seeders []SeederContract
}

func NewDatabaseSeeder(db *gorm.DB, seeders ...SeederContract) *DatabaseSeeder {
	dbs := new(DatabaseSeeder)
	dbs.db = db
	if len(seeders) > 0 {
		dbs.seeders = seeders
	}
	return dbs
}

// Run runs the database seeder
func (dbs *DatabaseSeeder) Run(c *fiber.Ctx) (err error) {
	// If no seeders provided, get default seeders and assign default seeders to dbs.seeders
	if len(dbs.seeders) == 0 {
		dbs.seeders = dbs.GetDefaultSeeders()
	}

	// Loop each dbs.seeder and run go routine function
	for _, seeder := range dbs.seeders {
		err = seeder.Run(c)
		if err != nil {
			return
		}
	}

	return
}

// GetDefaultSeeders returns the default seeders
func (dbs *DatabaseSeeder) GetDefaultSeeders() []SeederContract {
	// Prepare seeder repositories dependencies
	var (
		userRepositoy         = repositories.NewUserRepository(dbs.db)
		userSettingRepository = repositories.NewUserSettingRepository(dbs.db)
		tokenRepository       = repositories.NewTokenRepository(dbs.db)
	)

	// return seeders
	return []SeederContract{
		NewUserSeeder(userRepositoy),
		NewUserSettingSeeder(userSettingRepository, userRepositoy),
		NewTokenSeeder(tokenRepository, userRepositoy),
	}
}
