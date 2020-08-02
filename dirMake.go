package util

import (
	"os"
	"path/filepath"
)

func DirMake(path string) error {
	var err error

	dir, _ := filepath.Split(path)

	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
	}

	return err
}
