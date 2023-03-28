package resources

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Bytes returns Response as bytes
func (response Response) Bytes() []byte {
	bytes, err := json.Marshal(response)
	if err != nil {
		return []byte{}
	}
	return bytes

}

// ResponseFromError makes/fills Response from the given error
func ResponseFromError(err *fiber.Error) Response {
	return Response{
		Code:    err.Code,
		Message: err.Message,
	}
}
