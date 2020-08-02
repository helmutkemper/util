package util

import "os"

func DirCheckExists(path string) (exists bool) {
	if info, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) == false && info.IsDir() == true {
			return true
		}
	}
	return false
}
