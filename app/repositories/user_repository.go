package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository instantiates a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// All returns all users in the database
func (repository *UserRepository) All(c *fiber.Ctx) (users []models.User, err error) {
	err = repository.db.Find(&users).Error
	return
}

// FindByEmail finds a user by email
func (repository *UserRepository) FindByEmail(c *fiber.Ctx, email string) (user models.User, err error) {
	err = repository.db.Where("email = ?", email).First(&user).Error
	return
}

// Insert inserts users into the database
func (repository *UserRepository) Insert(c *fiber.Ctx, users ...*models.User) (int64, error) {
	for _, user := range users {
		if user.Id == uuid.Nil {
			user.Id = uuid.New()
		}
		user.CreatedAt = time.Now()
	}
	result := repository.db.Create(users)
	return result.RowsAffected, result.Error
}

// UpdateById updates model in the database by given id
func (repository *UserRepository) UpdateById(c *fiber.Ctx, user *models.User) (int64, error) {
	// refresh model updated at
	user.UpdatedAt = null.NewTime(time.Now(), true)
	// update
	result := repository.db.Model(user).Updates(user)
	return result.RowsAffected, result.Error
}
