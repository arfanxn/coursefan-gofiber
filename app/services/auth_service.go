package services

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/jwth"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/mailh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/reflecth"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepository  *repositories.UserRepository
	mediaRepository *repositories.MediaRepository
	tokenRepository *repositories.TokenRepository
}

// NewAuthService instantiates a new AuthService
func NewAuthService(
	userRepository *repositories.UserRepository,
	mediaRepository *repositories.MediaRepository,
	tokenRepository *repositories.TokenRepository,
) *AuthService {
	return &AuthService{
		userRepository:  userRepository,
		mediaRepository: mediaRepository,
		tokenRepository: tokenRepository,
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

	// validate whether email already registered
	if userMdl, _ := service.userRepository.FindByEmail(c, input.Email); userMdl.Id != uuid.Nil {
		return data, exceptions.NewValidationError("email",
			fmt.Sprintf(`email %s has already been taken.`, input.Email))
	}

	var (
		userMdl = models.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: string(password),
		}
		avatarMediaMdl models.Media
	)

	_, err = service.userRepository.Insert(c, &userMdl)
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
		_, err = service.mediaRepository.Insert(c, &avatarMediaMdl)
		if err != nil {
			return data, err
		}
		data.Avatar = resources.NewMediaFromModel(avatarMediaMdl)
	}

	return data, nil
}

// Register
func (service *AuthService) ForgotPassword(c *fiber.Ctx, input requests.AuthForgotPassword) (
	err error) {
	var (
		userMdl  models.User
		tokenMdl models.Token
	)

	// User relateds
	userMdl, err = service.userRepository.FindByEmail(c, input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // return not found if not found
			err = exceptions.NewValidationError("email", "No user found with email "+input.Email)
		}
		return
	}

	// Token relateds
	tokenMdl, err = service.tokenRepository.FindByTokenableAndType(
		c,
		reflecth.GetTypeName(userMdl),
		userMdl.Id.String(),
		enums.TokenTypeResetPassword,
	)
	// if token not found or has been used or expired then  and create a new one
	if errors.Is(err, gorm.ErrRecordNotFound) || (tokenMdl.IsUsed()) || tokenMdl.IsExpired() {
		tokenMdl.TokenableType = reflecth.GetTypeName(userMdl)
		tokenMdl.TokenableId = userMdl.Id
		tokenMdl.Type = enums.TokenTypeResetPassword
		tokenMdl.UsedAt = sql.NullTime{Time: time.Time{}, Valid: false}
		tokenMdl.ExpiredAt = sql.NullTime{Time: time.Now().Add(time.Hour / 2), Valid: true} // give 30 mins expiration
		tokenMdl.GenerateBody(models.TokenBodyNumeric, 6)
		_, err = service.tokenRepository.Save(c, &tokenMdl)
		if err != nil {
			return
		}
	}

	// Send Token to User's email
	err = mailh.Send(os.Getenv("MAIL_SENDER"),
		"OTP | Reset Password",
		"Your reset password OTP is "+tokenMdl.Body,
		input.Email,
	)
	if err != nil {
		return
	}

	return nil
}
