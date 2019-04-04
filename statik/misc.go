//go:generate rm -f statik.go
//go:generate statik -f -src=. -dest=..

package statik

import (
	"errors"
	"io/ioutil"

	"github.com/rakyll/statik/fs"
)

// ReadFile reads a file content from the embedded filesystem.
func ReadFile(name string) ([]byte, error) {
	fs, err := fs.New()
	if err != nil {
		return nil, errors.New("failed to open statik filesystem")
	}

	file, err := fs.Open(name)
	if err != nil {
		return nil, errors.New("failed to open statik file")
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("failed to read statik file")
	}

	return data, nil
}
