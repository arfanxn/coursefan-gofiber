package repositories

import (
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/gofiber/fiber/v2"
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
	result := repository.db.Where("email = ?", email).First(&user)
	err = result.Error
	return
}
