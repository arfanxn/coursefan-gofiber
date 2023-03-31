package validatorp

import (
	"github.com/arfanxn/coursefan-gofiber/bootstrap"
	"github.com/go-playground/validator/v10"
)

var (
	// validator.Validate instance
	validatorInstance *validator.Validate = nil
)

// GetValidate returns a singleton of validator.Validate and error if there an error
func GetValidator() (*validator.Validate, error) {
	if validatorInstance != nil {
		return validatorInstance, nil
	}
	var err error
	validatorInstance, err = bootstrap.NewValidator()
	return validatorInstance, err
}
