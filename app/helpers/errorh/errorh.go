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
