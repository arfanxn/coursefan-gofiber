package controllers

import (
	"os"
	"strconv"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/validationh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service *services.AuthService
}

// NewAuthController instantiates a new AuthController
func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (controller *AuthController) Login(c *fiber.Ctx) (err error) {
	var input requests.AuthLogin
	c.BodyParser(&input)
	if validationErrs := validationh.ValidateStruct(input, "en"); validationErrs != nil {
		response := resources.NewResponseValidationErrs(validationErrs)
		return c.Send(response.Bytes())
	}

	data, err := controller.service.Login(c, input)
	if err != nil {
		return err
	}

	// Get auth expiration seconds from environment variable
	authExpSec, err := strconv.ParseInt(os.Getenv("AUTH_EXP"), 10, 64)
	if err != nil {
		return err
	}
	// Set token to the cookie
	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("AUTH_COOKIE_NAME"),
		Path:     "/",
		Value:    data.AccessToken,
		HTTPOnly: true,
		MaxAge:   int(authExpSec),
	})

	return c.Send(resources.Response{
		Code:    fiber.StatusOK,
		Message: "Login successfully",
		Data:    data,
	}.Bytes())
}
