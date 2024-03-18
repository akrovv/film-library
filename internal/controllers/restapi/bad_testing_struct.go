package restapi

import (
	"errors"
)

type BadReader struct {
}

func (br BadReader) Read(p []byte) (int, error) {
	return 0, errors.New("some error")
}
