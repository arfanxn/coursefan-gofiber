package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type WalletSeeder struct {
	repository     *repositories.WalletRepository
	userRepository *repositories.UserRepository
}

// NewWalletSeeder instantiates a new WalletSeeder
func NewWalletSeeder(
	repository *repositories.WalletRepository,
	userRepository *repositories.UserRepository,
) *WalletSeeder {
	return &WalletSeeder{
		repository:     repository,
		userRepository: userRepository,
	}
}

// Run runs the seeder
func (seeder *WalletSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all users
	users, err := seeder.userRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var wallets []*models.Wallet
	for _, user := range users {
		wallet := factories.FakeWallet()
		wallet.OwnerId = user.Id
		wallets = append(wallets, &wallet)
	}
	_, err = seeder.repository.Insert(c, wallets...)

	return
}
