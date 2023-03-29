package fileh

import (
	"io"
	"os"
	"path"
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

// Save saves a file by given path
func Save(dstPath string, file io.Reader) (err error) {
	errMkdir := os.MkdirAll(path.Dir(dstPath), os.ModePerm)
	if errMkdir != nil {
		err = errMkdir
		return
	}
	fileDst, errCreate := os.Create(dstPath)
	if errCreate != nil {
		err = errCreate
		return
	}
	defer fileDst.Close()

	_, errCopy := io.Copy(fileDst, file)
	if errCopy != nil {
		err = errCopy
		return
	}

	return
}
