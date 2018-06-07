package util

import "os"

/*
func check(){
  if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
    // path/to/whatever does not exist
  }

  if _, err := os.Stat("/path/to/whatever"); err == nil {
    // path/to/whatever exists
  }
}
*/

func CheckFileExists(name string) bool {
  if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
      return false
    }
  }
  return true
}