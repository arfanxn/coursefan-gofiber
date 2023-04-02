package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository instantiates a new TransactionRepository
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// All returns all transactions in the database
func (repository *TransactionRepository) All(c *fiber.Ctx) (transactions []models.Transaction, err error) {
	err = repository.db.Find(&transactions).Error
	return
}

// Find finds model by id
func (repository *TransactionRepository) Find(c *fiber.Ctx, id string) (transaction models.Transaction, err error) {
	err = repository.db.Where("id = ?", id).First(&transaction).Error
	return
}

// Insert inserts models into the database
func (repository *TransactionRepository) Insert(c *fiber.Ctx, transactions ...*models.Transaction) (int64, error) {
	for _, transaction := range transactions {
		if transaction.Id == uuid.Nil {
			transaction.Id = uuid.New()
		}
		transaction.CreatedAt = time.Now()
	}
	result := repository.db.Create(transactions)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *TransactionRepository) UpdateById(c *fiber.Ctx, transaction *models.Transaction) (int64, error) {
	// refresh model updated at
	transaction.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(transaction).Updates(transaction)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *TransactionRepository) DeleteByIds(c *fiber.Ctx, transactions ...*models.Transaction) (int64, error) {
	result := repository.db.Delete(transactions)
	return result.RowsAffected, result.Error
}
