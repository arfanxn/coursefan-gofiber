package services

import (
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
	"github.com/arfanxn/coursefan-gofiber/app/helpers/synch"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	data.UpdatedAt = userMdl.UpdatedAt
	// if "AUTH_RETURN_TOKEN" set to true, return token on response body after successful login
	if isTrue, err := strconv.ParseBool(os.Getenv("AUTH_RETURN_TOKEN")); isTrue && (err == nil) {
		data.AccessToken = token
	}
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
	// if token not found or has been used or expired then  create a new one
	if errors.Is(err, gorm.ErrRecordNotFound) || (tokenMdl.IsUsed()) || tokenMdl.IsExpired() {
		tokenMdl.TokenableType = reflecth.GetTypeName(userMdl)
		tokenMdl.TokenableId = userMdl.Id
		tokenMdl.Type = enums.TokenTypeResetPassword
		tokenMdl.UsedAt = null.NewTime(time.Time{}, false)
		tokenMdl.ExpiredAt = null.NewTime(time.Now().Add(time.Hour/2), true) // give 30 mins expiration
		tokenMdl.BodyGenerate(models.TokenBodyNumeric, 6)
		affected, err := service.tokenRepository.Save(c, &tokenMdl)
		logrus.Info(fmt.Sprintf("Affected %v", affected))
		if err != nil {
			return err
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

// ResetPassword resets User's password if the given otp is valid
func (service *AuthService) ResetPassword(c *fiber.Ctx, input requests.AuthResetPassword) (err error) {
	var (
		syncronizer = synch.NewSyncronizer()
		userMdl     models.User
		tokenMdl    models.Token
	)
	defer syncronizer.Close()
	userMdl, err = service.userRepository.FindByEmail(c, input.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = exceptions.NewValidationError("email", "No user found with email "+input.Email)
		return
	} else if err != nil {
		return
	}
	tokenMdl, err = service.tokenRepository.FindByTokenableAndType(
		c,
		reflecth.GetTypeName(userMdl),
		userMdl.Id.String(),
		enums.TokenTypeResetPassword,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) ||
		tokenMdl.IsUsed() ||
		tokenMdl.IsExpired() ||
		(tokenMdl.Body != input.Otp) {
		err = exceptions.NewValidationError("otp", "OTP has been used or expired or invalid")
		return
	} else if err != nil {
		return
	}
	syncronizer.WG().Add(2)
	go func() { // goroutine for update token
		defer syncronizer.WG().Done()
		if syncronizer.Err() != nil {
			return
		}
		syncronizer.M().Lock()
		tokenMdl.UsedAt = null.NewTime(time.Now(), true)
		_, err := service.tokenRepository.UpdateById(c, &tokenMdl)
		syncronizer.M().Unlock()
		if err != nil {
			syncronizer.Err(err)
			return
		}
	}()
	go func() { // goroutine for update user
		defer syncronizer.WG().Done()
		if syncronizer.Err() != nil {
			return
		}
		passwordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			syncronizer.Err(err)
			return
		}
		syncronizer.M().Lock()
		userMdl.Password = string(passwordBytes)
		_, err = service.userRepository.UpdateById(c, &userMdl)
		syncronizer.M().Unlock()
		if err != nil {
			syncronizer.Err(err)
		}
	}()
	syncronizer.WG().Wait()
	if err != nil {
		return
	}
	if err = syncronizer.Err(); err != nil {
		return
	}
	return
}
