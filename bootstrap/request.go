package bootstrap

import (
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/ggicci/httpin"
)

// Request will bootstraping aplication request relateds
func Request() error {
	httpin.RegisterNamedDecoder("query_filters", httpin.ValueTypeDecoderFunc(requests.Query{}.DecodeFilters))
	return nil
}
