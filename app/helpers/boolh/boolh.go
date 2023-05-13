package boolh

// Ternary
func Ternary[T any](
	condition bool,
	onTrue T,
	onFalse T,
) T {
	if condition {
		return onTrue
	} else {
		return onFalse
	}
}

// ToPointer converts the given boolean value to a pointer of boolean value
func ToPointer(value bool) *bool {
	return &value
}
