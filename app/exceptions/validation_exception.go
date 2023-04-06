package exceptions

/*
 * ValidationErrors (wrapper for ValidationError)
 */

type ValidationErrors struct {
	Errors []*ValidationError
}

func NewValidationErrors(errs []*ValidationError) *ValidationErrors {
	validationErrs := new(ValidationErrors)
	validationErrs.Errors = errs
	return validationErrs
}

func (validationErrs *ValidationErrors) Error() string {
	if len(validationErrs.Errors) == 0 {
		return ""
	}
	return validationErrs.Errors[0].Error()
}

/*
 * ValidationError
 */

type ValidationError struct {
	Field, Message string
}

func (validationErr *ValidationError) Error() string {
	return validationErr.Message
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}
