// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/controllers"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"gorm.io/gorm"
)

// Injectors from injector.go:

func AuthController(db *gorm.DB) *controllers.AuthController {
	userRepository := repositories.NewUserRepository(db)
	mediaRepository := repositories.NewMediaRepository(db)
	tokenRepository := repositories.NewTokenRepository(db)
	authService := services.NewAuthService(userRepository, mediaRepository, tokenRepository)
	authController := controllers.NewAuthController(authService)
	return authController
}
