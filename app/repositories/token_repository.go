package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

// NewTokenRepository instantiates a new TokenRepository
func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

// Find finds a token by id
func (repository *TokenRepository) Find(c *fiber.Ctx, id string) (token models.Token, err error) {
	err = repository.db.Where("id = ?", id).First(&token).Error
	return
}

// FindByTokenableAndType finds a token by tokenable and type
func (repository *TokenRepository) FindByTokenableAndType(
	c *fiber.Ctx, tokenableTyp string, tokenableId string, typ string,
) (token models.Token, err error) {
	err = repository.db.
		Where("tokenable_type = ?", tokenableTyp).
		Where("tokenable_id = ?", tokenableId).
		Where("type = ?", typ).
		First(&token).Error
	return
}

// Insert inserts tokens into the database
func (repository *TokenRepository) Insert(c *fiber.Ctx, tokens ...*models.Token) (int64, error) {
	for _, token := range tokens {
		if token.Id == uuid.Nil {
			token.Id = uuid.New()
		}
		token.CreatedAt = time.Now()
	}
	result := repository.db.Create(tokens)
	return result.RowsAffected, result.Error
}

// Save updates value in database. If value doesn't contain a matching primary key, value is inserted.
func (repository *TokenRepository) Save(c *fiber.Ctx, token *models.Token) (int64, error) {
	if token.Id == uuid.Nil {
		token.Id = uuid.New()
	}
	// if model is already created then update the model updated_at
	if token.CreatedAt != (time.Time{}) {
		token.UpdatedAt = null.NewTime(time.Now(), true)
	} else { // otherwise
		token.CreatedAt = time.Now()
	}
	result := repository.db.Model(token).Save(token)
	return result.RowsAffected, result.Error
}

// UpdateById
func (repository *TokenRepository) UpdateById(c *fiber.Ctx, token *models.Token) (int64, error) {
	result := repository.db.Model(token).Updates(token)
	return result.RowsAffected, result.Error
}

// DeleteByIds deletes the entities associated with the given ids
func (repository *TokenRepository) DeleteByIds(c *fiber.Ctx, tokens ...*models.Token) (int64, error) {
	result := repository.db.Delete(tokens)
	return result.RowsAffected, result.Error
}
