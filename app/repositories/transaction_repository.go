package repositories

import (
	"fmt"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
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

// All returns all reviews in the database
func (repository *TransactionRepository) All(c *fiber.Ctx, queries ...requests.Query) (reviews []models.Transaction, err error) {
	tx := repository.db
	if query := sliceh.FirstOrNil(queries); query != nil {
		tx = gormh.BuildFromRequestQuery(repository.db, models.Transaction{}, *query)
	}
	err = tx.Find(&reviews).Error
	return
}

// AllByAuthUserWallet returns all transactions
func (repository *TransactionRepository) AllByAuthUserWallet(c *fiber.Ctx, query requests.Query) (
	reviews []models.Transaction, err error) {
	err = gormh.BuildFromRequestQuery(repository.db, models.Transaction{}, query).
		Joins(
			fmt.Sprintf("JOIN %s ON %s.%s = %s.%s OR %s.%s = %s.%s",
				models.Wallet{}.TableName(),
				// ----
				models.Wallet{}.TableName(),
				"id",
				models.Transaction{}.TableName(),
				"receiver_id",
				// ----
				models.Wallet{}.TableName(),
				"id",
				models.Transaction{}.TableName(),
				"sender_id",
			)).
		Joins(
			fmt.Sprintf("JOIN %s ON %s.%s = %s.%s",
				models.User{}.TableName(),
				models.User{}.TableName(),
				"id",
				models.Wallet{}.TableName(),
				"owner_id",
			)).
		Where(models.User{}.TableName()+".id = ?", ctxh.MustGetUser(c).Id).
		Distinct().Find(&reviews).Error
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
