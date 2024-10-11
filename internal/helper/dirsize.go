package dirsize

import (
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
)

func DirSize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, "DirSize")
		}

		if !info.IsDir() {
			size += info.Size()
		}

		return errors.Wrap(err, "DirSize")
	})

	return size, errors.Wrap(err, "DirSize")
}
