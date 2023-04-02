package seeders

import "github.com/gofiber/fiber/v2"

type SeederContract interface {
	Run(*fiber.Ctx) error
}
