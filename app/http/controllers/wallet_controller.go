package controllers

import (
	"github.com/arfanxn/coursefan-gofiber/app/enums"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/responseh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/policies"
	"github.com/arfanxn/coursefan-gofiber/app/services"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type WalletController struct {
	policy  *policies.WalletPolicy
	service *services.WalletService
}

func NewWalletController(
	policy *policies.WalletPolicy,
	service *services.WalletService,
) *WalletController {
	return &WalletController{
		policy:  policy,
		service: service,
	}
}

// FindByAuthUser wallet by the current logged in user
func (controller *WalletController) FindByAuthUser(c *fiber.Ctx) (err error) {
	input := requests.Query{}
	err = input.FromContext(c)
	if err != nil {
		return
	}
	input.AddFilter(requests.QueryFilter{
		Column:   models.Wallet{}.TableName() + ".owner_id",
		Operator: enums.QueryFilterOperatorEquals,
		Values:   []any{ctxh.MustGetUser(c).Id},
	})
	// No api policy required, only the authenticated user can access their own wallet
	data, err := controller.service.Find(c, input)
	if err != nil {
		return err
	}
	return responseh.Write(c, resources.Response{
		Code:    fiber.StatusOK,
		Message: "Successfully retrieve wallet",
		Data:    data,
	})
}
