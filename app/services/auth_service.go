package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/jwth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository  *repositories.UserRepository
	mediaRepository *repositories.MediaRepository
}

// NewAuthService instantiates a new AuthService
func NewAuthService(
	userRepository *repositories.UserRepository,
	mediaRepository *repositories.MediaRepository,
) *AuthService {
	return &AuthService{
		userRepository:  userRepository,
		mediaRepository: mediaRepository,
	}
}

// Login
func (service *AuthService) Login(c *fiber.Ctx, input requests.AuthLogin) (data resources.AuthLogin, err error) {
	// Get auth expiration in seconds from environment variable
	authExpSec, err := strconv.ParseInt(os.Getenv("AUTH_EXP"), 10, 64)
	if err != nil {
		return
	}

	userMdl, err := service.userRepository.FindByEmail(c, input.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = exceptions.AuthCredentialsDoesNotMatch
		return
	} else if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userMdl.Password), []byte(input.Password))
	if err != nil {
		err = exceptions.AuthCredentialsDoesNotMatch
		return
	}

	token, err := jwth.Encode(os.Getenv("APP_KEY"), map[string]any{
		"authorized": true,
		"user":       any(userMdl),
		"exp":        time.Now().Add(time.Minute * time.Duration(authExpSec)).Unix(),
	})
	if err != nil {
		return
	}

	data.Id = userMdl.Id.String()
	data.Name = userMdl.Name
	data.Email = userMdl.Email
	data.CreatedAt = userMdl.CreatedAt
	data.UpdatedAt = null.NewTime(userMdl.UpdatedAt.Time, userMdl.UpdatedAt.Valid)
	data.AccessToken = token
	return
}

// Register
func (service *AuthService) Register(c *fiber.Ctx, input requests.AuthRegister) (
	data resources.AuthRegister, err error) {
	// Hash the user password
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	var (
		userMdl = models.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: string(password),
		}
		avatarMediaMdl models.Media
	)

	service.userRepository.Insert(c, &userMdl)
	if err != nil {
		return
	}
	data.Id = userMdl.Id.String()
	data.Name = userMdl.Name
	data.Email = userMdl.Email
	data.CreatedAt = userMdl.CreatedAt
	data.UpdatedAt = null.NewTime(userMdl.UpdatedAt.Time, userMdl.UpdatedAt.Valid)

	// if user's avatar provided then save it
	if input.Avatar != nil {
		avatarMediaMdl.ModelType = reflecth.GetTypeName(userMdl)
		avatarMediaMdl.ModelId = userMdl.Id
		avatarMediaMdl.SetFileHeader(input.Avatar)
		err = service.mediaRepository.Insert(c, &avatarMediaMdl)
		if err != nil {
			return data, err
		}
		data.Avatar = resources.NewMediaFromModel(avatarMediaMdl)
	}

	return data, nil
}
