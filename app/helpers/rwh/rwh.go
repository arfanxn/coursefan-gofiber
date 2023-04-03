package rwh

import (
	"bytes"
	"io"

	"github.com/clarketm/json"
)

func StructToReadCloser(v any) (readerCloser io.ReadCloser, err error) {
	marshalled, err := json.Marshal(v)
	if err != nil {
		return
	}
	reader := bytes.NewReader(marshalled)
	readerCloser = io.NopCloser(reader)
	return
}
