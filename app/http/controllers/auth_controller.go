package controllers

import (
	"os"
	"strconv"
	"time"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/validateh"
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

// Login
func (controller *AuthController) Login(c *fiber.Ctx) (err error) {
	var input requests.AuthLogin
	c.BodyParser(&input)
	if validationErrs := validateh.ValidateStruct(input, ctxh.GetAcceptLang(c)); validationErrs != nil {
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
		Expires:  time.Now().Add(time.Duration(authExpSec) * time.Second),
	})

	resBody := resources.Response{
		Code:    fiber.StatusOK,
		Message: "Login successfully",
		Data:    data,
	}
	return c.Status(resBody.Code).Send(resBody.Bytes())
}

// Logout will signing out the signed in if user
func (controller *AuthController) Logout(c *fiber.Ctx) error {
	// Delete token from cookie
	c.Cookie(&fiber.Cookie{
		Name:     os.Getenv("AUTH_COOKIE_NAME"),
		Path:     "/",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
		Expires:  time.Now().Add(time.Second),
	})
	resBody := resources.Response{
		Code:    fiber.StatusOK,
		Message: "Logout successfully",
	}
	return c.Status(resBody.Code).Send(resBody.Bytes())
}

// Register
func (controller *AuthController) Register(c *fiber.Ctx) (err error) {
	var input requests.AuthRegister
	c.BodyParser(&input)
	input.Avatar = ctxh.GetFileHeader(c, "avatar")
	if validationErrs := validateh.ValidateStruct(input, ctxh.GetAcceptLang(c)); validationErrs != nil {
		response := resources.NewResponseValidationErrs(validationErrs)
		return c.Send(response.Bytes())
	}

	data, err := controller.service.Register(c, input)
	if err != nil {
		return err
	}

	resBody := resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Register successfully",
		Data:    data,
	}
	return c.Status(resBody.Code).Send(resBody.Bytes())
}

// ForgotPassword
func (controller *AuthController) ForgotPassword(c *fiber.Ctx) (err error) {
	var input requests.AuthForgotPassword
	c.BodyParser(&input)
	if validationErrs := validateh.ValidateStruct(input, ctxh.GetAcceptLang(c)); validationErrs != nil {
		response := resources.NewResponseValidationErrs(validationErrs)
		return c.Send(response.Bytes())
	}

	err = controller.service.ForgotPassword(c, input)
	if err != nil {
		return err
	}

	resBody := resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully sent reset password token to " + input.Email,
	}
	return c.Status(resBody.Code).Send(resBody.Bytes())
}

// ResetPassword
func (controller *AuthController) ResetPassword(c *fiber.Ctx) (err error) {
	var input requests.AuthResetPassword
	c.BodyParser(&input)
	if validationErrs := validateh.ValidateStruct(input, ctxh.GetAcceptLang(c)); validationErrs != nil {
		response := resources.NewResponseValidationErrs(validationErrs)
		return c.Send(response.Bytes())
	}

	err = controller.service.ResetPassword(c, input)
	if err != nil {
		return err
	}

	resBody := resources.Response{
		Code:    fiber.StatusCreated,
		Message: "Successfully reset password",
	}
	return c.Status(resBody.Code).Send(resBody.Bytes())
}
