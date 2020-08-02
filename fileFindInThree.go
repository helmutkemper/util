package util

import (
	"errors"
	"os"
	"path/filepath"
)

func FileFindInThree(name string) (filePath string, err error) {
	var fileInfo os.FileInfo
	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if name == info.Name() {
				filePath = path
				fileInfo = info
				return nil
			}

			return nil
		},
	)

	if err != nil {
		return
	}

	if filePath == "" || fileInfo.IsDir() == true {
		filePath = ""
		err = errors.New("file not found")
	}

	return
}
