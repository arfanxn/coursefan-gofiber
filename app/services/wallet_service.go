package services

import (
	"net/url"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/ctxh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/gofiber/fiber/v2"
)

type WalletService struct {
	repository *repositories.WalletRepository
}

func NewWalletService(repository *repositories.WalletRepository) *WalletService {
	return &WalletService{
		repository: repository,
	}
}

// All
func (service *WalletService) All(c *fiber.Ctx, input requests.Query) (
	pagination resources.Pagination[resources.Wallet], err error) {
	walletMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	}
	walletRess := sliceh.Map(walletMdls, func(walletMdl models.Wallet) resources.Wallet {
		walletRes := resources.Wallet{}
		walletRes.FromModel(walletMdl)
		return walletRes
	})
	pagination.SetItems(walletRess)
	pagination.SetPageFromOffsetLimit(int64(input.Offset), int(input.Limit.Int64))
	pagination.SetURL(errorh.Must(url.Parse(ctxh.GetFullURIString(c))))
	return
}

// Find
func (service *WalletService) Find(c *fiber.Ctx, input requests.Query) (
	data resources.Wallet, err error) {
	walletMdls, err := service.repository.All(c, input)
	if err != nil {
		return
	} else if len(walletMdls) == 0 {
		err = fiber.ErrNotFound
		return
	}
	data.FromModel(walletMdls[0])
	return
}
