package seeders

import (
	"math/rand"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/gofiber/fiber/v2"
)

type TransactionSeeder struct {
	repository       *repositories.TransactionRepository
	courseRepository *repositories.CourseRepository
	walletRepository *repositories.WalletRepository
}

// NewTransactionSeeder instantiates a new TransactionSeeder
func NewTransactionSeeder(
	repository *repositories.TransactionRepository,
	courseRepository *repositories.CourseRepository,
	walletRepository *repositories.WalletRepository,
) *TransactionSeeder {
	return &TransactionSeeder{
		repository:       repository,
		courseRepository: courseRepository,
		walletRepository: walletRepository,
	}
}

// Run runs the seeder
func (seeder *TransactionSeeder) Run(c *fiber.Ctx) (err error) {
	// Get all wallets
	wallets, err := seeder.walletRepository.All(c)
	if err != nil {
		return
	}
	// Get all courses
	courses, err := seeder.courseRepository.All(c)
	if err != nil {
		return
	}
	// Seed
	var transactions []*models.Transaction
	for _, wallet := range wallets {
		course := courses[rand.Intn(len(courses))-1]
		transaction := factories.FakeTransaction()
		transaction.TransactionableType = reflecth.GetTypeName(course)
		transaction.TransactionableId = course.Id
		transaction.SenderId = wallet.Id
		transaction.ReceiverId = wallets[0].Id
		transactions = append(transactions, &transaction)
	}
	_, err = seeder.repository.Insert(c, transactions...)

	return
}
