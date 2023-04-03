package errorh

import (
	"errors"

	"gorm.io/gorm"
)

// IsGormErrRecordNotFound cheks whether the given err is an error of gorm.ErrRecordNotFound, return true if it is
func IsGormErrRecordNotFound(errAny any) bool {
	if err, ok := errAny.(error); ok {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
	return false
}

// Must returns value of first argument and if second argument is not nil then it will panic
func Must[T any](myvar T, err error) T {
	if err != nil {
		panic(err)
	}
	return myvar
}
