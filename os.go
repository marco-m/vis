package vis

import (
	"errors"
	"os"
)

func RemoveAllFiles(files ...string) error {
	var errs []error
	for _, f := range files {
		errs = append(errs, os.Remove(f))
	}
	// FIXME should filter and return only "real" errors
	return errors.Join(errs...)
}
