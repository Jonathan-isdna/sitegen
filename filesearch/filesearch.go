package filesearch

import (
  "path/filepath"
  "os"
  "github.com/Jonathan-isdna/sitegen/gf"
)

func visit(path string, f os.FileInfo, err error) error {
  // fmt.Printf("Visited: %s\n", path)
  if f.IsDir() {
    return nil
  } else {

  }
  return nil
}

func Search(startPath string) []string {
  root := startPath
  pathList := make([]string, 0)
  err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
    if !f.IsDir() {
      pathList = append(pathList, path)
    }
    return err
  })
  gf.Check(err)
  return pathList
}
