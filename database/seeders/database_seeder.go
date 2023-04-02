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
		dbs.seeders = dbs.DefaultSeeders()
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

// DefaultSeeders returns the default seeders
func (dbs *DatabaseSeeder) DefaultSeeders() []SeederContract {
	// Prepare seeder repositories dependencies
	var (
		userRepositoy          = repositories.NewUserRepository(dbs.db)
		userProfileRepository  = repositories.NewUserProfileRepository(dbs.db)
		userSettingRepository  = repositories.NewUserSettingRepository(dbs.db)
		tokenRepository        = repositories.NewTokenRepository(dbs.db)
		notificationRepository = repositories.NewNotificationRepository(dbs.db)
		messageRepository      = repositories.NewMessageRepository(dbs.db)
		walletRepository       = repositories.NewWalletRepository(dbs.db)
		transactionRepository  = repositories.NewTransactionRepository(dbs.db)
		courseRepository       = repositories.NewCourseRepository(dbs.db)
	)

	// return seeders
	return []SeederContract{
		NewUserSeeder(userRepositoy),
		NewUserProfileSeeder(userProfileRepository, userRepositoy),
		NewUserSettingSeeder(userSettingRepository, userRepositoy),
		NewTokenSeeder(tokenRepository, userRepositoy),
		NewNotificationSeeder(notificationRepository, userRepositoy),
		NewMessageSeeder(messageRepository, userRepositoy),
		NewWalletSeeder(walletRepository, userRepositoy),
		NewTransactionSeeder(transactionRepository, courseRepository, walletRepository),
	}
}
