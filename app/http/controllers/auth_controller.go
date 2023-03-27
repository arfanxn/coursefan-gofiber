package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/jwth"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	conn "github.com/arfanxn/coursefan-gofiber/database/connection"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

// NewAuthController instantiates a new AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	// Get cookie max age from config env variable
	authExpSec, err := strconv.ParseInt(os.Getenv("AUTH_EXP"), 10, 64)
	if err != nil {
		return err
	}

	db, err := conn.GetGORM()
	if err != nil {
		return err
	}
	var user models.User
	db.First(&user)

	token, err := jwth.Encode(os.Getenv("APP_KEY"), map[string]any{
		"authorized": true,
		"user":       user,
		"exp":        time.Now().Add(time.Minute * time.Duration(authExpSec)).Unix(),
	})
	if err != nil {
		return err
	}

	// Set token to the cookie
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Path:     "/",
		Value:    token,
		HTTPOnly: true,
		MaxAge:   int(authExpSec),
	})

	return c.Send(resources.Response{
		Code:    fiber.StatusOK,
		Message: "Login successfully",
	}.Bytes())
}
