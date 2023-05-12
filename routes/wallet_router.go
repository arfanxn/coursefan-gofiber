package routes

import (
	"github.com/arfanxn/coursefan-gofiber/app/providers/controllerp"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/gofiber/fiber/v2"
)

// registerWalletRouter registers wallet module routes into the router
func registerWalletRouter(router fiber.Router) {
	walletController := controllerp.InitWalletController(databasep.MustGetGormDB())
	router.Get("/users/self/wallet", walletController.FindByAuthUser)
}
