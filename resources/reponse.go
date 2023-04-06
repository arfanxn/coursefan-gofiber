package resources

import (
	"github.com/arfanxn/coursefan-gofiber/app/exceptions"
	"github.com/clarketm/json"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code       int               `json:"code"`
	Message    string            `json:"message"`
	Errors     map[string]string `json:"errors,omitempty"`
	Data       any               `json:"data,omitempty"`
	Pagination any               `json:"pagination,omitempty"`
}

// Bytes returns Response as bytes
func (response Response) Bytes() []byte {
	bytes, err := json.Marshal(response)
	if err != nil {
		return []byte{}
	}
	return bytes

}

// FromError fills response from the given error
func (response *Response) FromError(err *fiber.Error) {
	response.Code = err.Code
	response.Message = err.Message
}

// FromValidationErrs fills response from the given validation errors
func (response *Response) FromValidationErrs(errs []*exceptions.ValidationError) {
	response.Code = fiber.StatusUnprocessableEntity
	response.Errors = map[string]string{}
	for index, err := range errs {
		if index == 0 {
			response.Message = err.Message
		}
		response.Errors[err.Field] = err.Message
	}
}
