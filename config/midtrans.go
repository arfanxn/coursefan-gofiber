package config

import (
	"os"

	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/midtrans/midtrans-go"
)

func ConfigureMidtrans() error {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")

	switch os.Getenv("MIDTRANS_ENVIRONMENT") {
	case "sandbox", "Sandbox", "SANDBOX":
		midtrans.Environment = midtrans.Sandbox
	case "production", "Production", "PRODUCTION":
		midtrans.Environment = midtrans.Production
	default:
		return exceptions.ENVVariableIsNotSetOrInvalid
	}

	return nil
}
