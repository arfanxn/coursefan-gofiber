package fileh

import (
	"os"
	"sync"
)

// BatchRemove removes many files/dirs in one
func BatchRemove(paths ...string) (err error) {
	wg := new(sync.WaitGroup)
	for _, path := range paths {
		if err != nil {
			return
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, path string) {
			defer wg.Done()
			err = os.RemoveAll(path)
		}(wg, path)
	}
	wg.Wait()

	return
}
