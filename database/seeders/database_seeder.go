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
		db                       = dbs.db
		permissionRepository     = repositories.NewPermissionRepository(db)
		roleRepository           = repositories.NewRoleRepository(db)
		permissionRoleRepository = repositories.NewPermissionRoleRepository(db)
		userRepository           = repositories.NewUserRepository(db)
		userProfileRepository    = repositories.NewUserProfileRepository(db)
		userSettingRepository    = repositories.NewUserSettingRepository(db)
		tokenRepository          = repositories.NewTokenRepository(db)
		notificationRepository   = repositories.NewNotificationRepository(db)
		messageRepository        = repositories.NewMessageRepository(db)
		walletRepository         = repositories.NewWalletRepository(db)
		transactionRepository    = repositories.NewTransactionRepository(db)
		courseRepository         = repositories.NewCourseRepository(db)
		cusRepository            = repositories.NewCourseUserRoleRepository(db)
		reviewRepository         = repositories.NewReviewRepository(db)
		lecturePartRepository    = repositories.NewLecturePartRepository(db)
		lectureRepository        = repositories.NewLectureRepository(db)
		progressRepository       = repositories.NewProgressRepository(db)
		progressUserRepository   = repositories.NewProgressUserRepository(db)
		discussionRepository     = repositories.NewDiscussionRepository(db)
		mediaRepository          = repositories.NewMediaRepository(db)
	)

	// return seeders
	return []SeederContract{
		NewPermissionSeeder(permissionRepository),
		NewRoleSeeder(roleRepository),
		NewPermissionRoleSeeder(permissionRoleRepository, permissionRepository, roleRepository),
		NewUserSeeder(userRepository),
		NewUserProfileSeeder(userProfileRepository, userRepository),
		NewUserSettingSeeder(userSettingRepository, userRepository),
		NewTokenSeeder(tokenRepository, userRepository),
		NewWalletSeeder(walletRepository, userRepository),
		NewNotificationSeeder(notificationRepository, userRepository),
		NewMessageSeeder(messageRepository, userRepository),
		NewCourseSeeder(courseRepository, userRepository),
		NewCourseUserRoleSeeder(cusRepository, courseRepository, userRepository, roleRepository),
		NewTransactionSeeder(transactionRepository, courseRepository, walletRepository),
		NewReviewSeeder(reviewRepository, cusRepository, roleRepository),
		NewLecturePartSeeder(lecturePartRepository, courseRepository),
		NewLectureSeeder(lectureRepository, lecturePartRepository),
		NewProgressSeeder(progressRepository, lectureRepository),
		NewProgressUserSeeder(progressUserRepository, progressRepository, userRepository),
		NewDiscussionSeeder(discussionRepository, lectureRepository, userRepository),
		NewMediaSeeder(mediaRepository, lectureRepository),
	}
}
