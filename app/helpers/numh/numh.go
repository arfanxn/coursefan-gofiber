// Numh stands for Number Helper
package numh

import (
	"math/rand"
	"time"
)

// MegabyteToByte converts Megabyte(s) to Byte(s)
func MegabyteToByte(mb int64) int64 {
	return mb * 1024 * 1024
}

// Random generates random numbers on range
func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
