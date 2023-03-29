package resources

import (
	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/clarketm/json"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
	Data    any               `json:"data,omitempty"`
}

// Bytes returns Response as bytes
func (response Response) Bytes() []byte {
	bytes, err := json.Marshal(response)
	if err != nil {
		return []byte{}
	}
	return bytes

}

// NewResponseError instantiates Response with the given error
func NewResponseError(err *fiber.Error) *Response {
	return &Response{
		Code:    err.Code,
		Message: err.Message,
	}
}

// NewResponseValidationErrs instantiates Response with the given validation errors
func NewResponseValidationErrs(errs []*exceptions.ValidationError) *Response {
	response := new(Response)
	response.Code = fiber.StatusUnprocessableEntity
	response.Errors = map[string]string{}
	for index, err := range errs {
		if index == 0 {
			response.Message = err.Message
		}
		response.Errors[err.Field] = err.Message
	}
	return response
}
