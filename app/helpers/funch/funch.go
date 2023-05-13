package funch

import "github.com/arfanxn/coursefan-gofiber/app/helpers/sliceh"

// Recursive makes a recursive call to the given function
func Recursive(fn func(*bool), maxTries ...int) {

	maxTry := sliceh.FirstOrNil(maxTries)
	var (
		repeat = true
		i      = 0
	)

	for repeat {
		if maxTry != nil && (*maxTry == i) {
			return
		}
		fn(&repeat)

		i++
	}
}
