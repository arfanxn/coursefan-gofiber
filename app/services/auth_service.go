package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/jwth"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository *repositories.UserRepository
}

// NewAuthService instantiates a new AuthService
func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (service *AuthService) Login(c *fiber.Ctx, input requests.AuthLogin) (res resources.AuthLogin, err error) {
	// Get auth expiration in seconds from environment variable
	authExpSec, err := strconv.ParseInt(os.Getenv("AUTH_EXP"), 10, 64)
	if err != nil {
		return
	}

	user, err := service.userRepository.FindByEmail(c, input.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = exceptions.AuthCredentialsDoesNotMatch
		return
	} else if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		err = exceptions.AuthCredentialsDoesNotMatch
		return
	}

	token, err := jwth.Encode(os.Getenv("APP_KEY"), map[string]any{
		"authorized": true,
		"user":       any(user),
		"exp":        time.Now().Add(time.Minute * time.Duration(authExpSec)).Unix(),
	})
	if err != nil {
		return
	}

	res.Id = user.Id.String()
	res.Name = user.Name
	res.Email = user.Email
	res.AccessToken = token
	return
}
