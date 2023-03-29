package repositories

import (
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository instantiates a new UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByEmail finds a user by email
func (repository *UserRepository) FindByEmail(c *fiber.Ctx, email string) (user models.User, err error) {
	err = repository.db.Where("email = ?", email).First(&user).Error
	return
}

// Insert inserts users into the database
func (repository *UserRepository) Insert(c *fiber.Ctx, users ...*models.User) (err error) {
	for _, user := range users {
		if user.Id == uuid.Nil {
			user.Id = uuid.New()
		}
		user.CreatedAt = time.Now()
	}
	err = repository.db.Create(users).Error
	return
}
