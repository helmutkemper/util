package util

import (
  "os"
  "path/filepath"
)

func DirMake( path string ) error {
  var err error

  if _, err = os.Stat( path ); os.IsNotExist( err ) {
    err = os.Mkdir( path, 0777 )
  }

  return err
}

func DirRemoveContents(dir string) error {
  d, err := os.Open(dir)
  if err != nil {
    return err
  }
  defer d.Close()
  names, err := d.Readdirnames(-1)
  if err != nil {
    return err
  }
  for _, name := range names {
    err = os.RemoveAll(filepath.Join(dir, name))
    if err != nil {
      return err
    }
  }
  return nil
}
