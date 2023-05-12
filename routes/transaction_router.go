package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerTransactionRouter registers transaction module routes into the router
func registerTransactionRouter(router fiber.Router) {
	transactionController := controllerp.InitTransactionController(databasep.MustGetGormDB())
	router.Get("/users/self/transactions", transactionController.AllByAuthUserWallet)
}
