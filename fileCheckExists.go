package util

import "os"

func FileCheckExists(name string) (exists bool) {
	if info, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) == false && info.IsDir() == false {
			return true
		}
	}
	return false
}
