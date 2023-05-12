package bootstrap

import (
	"github.com/arfanxn/coursefan-gofiber/config"
)

// Midtrans bootstraps midtrans library/package and configure it by the specified configuration on config file and environment variables
func Midtrans() (err error) {
	err = config.ConfigureMidtrans()
	return
}
