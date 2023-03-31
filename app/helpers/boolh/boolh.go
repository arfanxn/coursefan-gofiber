package boolh

// Ternary
func Ternary[T any](
	condition bool,
	onTrue T,
	onFalse T,
) any {
	if condition {
		return onTrue
	} else {
		return onFalse
	}
}
