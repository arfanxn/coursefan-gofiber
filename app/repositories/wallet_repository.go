package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/gormh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

// NewWalletRepository instantiates a new WalletRepository
func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

// All returns all wallets in the database
func (repository *WalletRepository) All(c *fiber.Ctx, queries ...requests.Query) (wallets []models.Wallet, err error) {
	tx := repository.db
	if query := sliceh.FirstOrNil(queries); query != nil {
		tx = gormh.BuildFromRequestQuery(repository.db, models.Wallet{}, *query)
	}
	err = tx.Find(&wallets).Error
	return
}

// FindById finds model by id
func (repository *WalletRepository) FindById(c *fiber.Ctx, id string) (wallet models.Wallet, err error) {
	err = repository.db.Where("id = ?", id).First(&wallet).Error
	return
}

// Insert inserts models into the database
func (repository *WalletRepository) Insert(c *fiber.Ctx, wallets ...*models.Wallet) (int64, error) {
	for _, wallet := range wallets {
		if wallet.Id == uuid.Nil {
			wallet.Id = uuid.New()
		}
		wallet.CreatedAt = time.Now()
	}
	result := repository.db.Create(wallets)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *WalletRepository) UpdateById(c *fiber.Ctx, wallet *models.Wallet) (int64, error) {
	// refresh model updated at
	wallet.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(wallet).Updates(wallet)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *WalletRepository) DeleteByIds(c *fiber.Ctx, wallets ...*models.Wallet) (int64, error) {
	result := repository.db.Delete(wallets)
	return result.RowsAffected, result.Error
}
