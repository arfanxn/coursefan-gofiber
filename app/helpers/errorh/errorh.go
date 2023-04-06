package errorh

import (
	"errors"

	"gorm.io/gorm"
)

// IsGormErrRecordNotFound checks whether the given err is an error of gorm.ErrRecordNotFound, return true if it is
func IsGormErrRecordNotFound(errAny any) bool {
	if err, ok := errAny.(error); ok {
		return errors.Is(err, gorm.ErrRecordNotFound)
	}
	return false
}

// IsGormError checks whether the given err is an errors from gorm package, return true if it is
func IsGormError(errAny any) bool {
	gormErrs := []error{
		gorm.ErrRecordNotFound,
		gorm.ErrInvalidTransaction,
		gorm.ErrNotImplemented,
		gorm.ErrMissingWhereClause,
		gorm.ErrUnsupportedRelation,
		gorm.ErrPrimaryKeyRequired,
		gorm.ErrModelValueRequired,
		gorm.ErrModelAccessibleFieldsRequired,
		gorm.ErrSubQueryRequired,
		gorm.ErrInvalidData,
		gorm.ErrUnsupportedDriver,
		gorm.ErrRegistered,
		gorm.ErrInvalidField,
		gorm.ErrEmptySlice,
		gorm.ErrDryRunModeUnsupported,
		gorm.ErrInvalidDB,
		gorm.ErrInvalidValue,
		gorm.ErrInvalidValueOfLength,
		gorm.ErrPreloadNotAllowed,
	}
	if err, ok := errAny.(error); ok {
		for _, gormErr := range gormErrs {
			if errors.Is(err, gormErr) {
				return true
			}
		}
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
