package fileio

import (
  "os"
  "io/ioutil"
  "bufio"
  "github.com/Jonathan-isdna/sitegen/gf"
  "path/filepath"
)

// Parses markdown file for title, date and markdown post
func FileReadLines(filename string) []string {
  // Open file into var vile and close it
  file, err := os.Open(filename)
  gf.Check(err)
  defer file.Close()

  // Create scanner and copy each line into a slice and return the slice
  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() { lines = append(lines, scanner.Text()) }
  gf.Check(scanner.Err())
  return lines
}

// Reads whole file and return as string
func FileRead(filename string) string {
  b, err := ioutil.ReadFile(filename)
  gf.Check(err)
  return string(b)
}

// Overwrites file with string input
func FileWrite(filename string, content string) {
  fileDir, _ := filepath.Split(filename)
  ensureDir(fileDir)
  byteString := []byte(content)
  err := ioutil.WriteFile(filename, byteString, os.ModePerm)
  gf.Check(err)
}

func ensureDir(dirPath string) {
  dirPath = "./" + dirPath
  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    os.MkdirAll(dirPath, os.ModePerm)
  }
}

func ResetBinFolder() {
  files, err := ioutil.ReadDir("bin")
  gf.Check(err)
  for _, file := range files {
    if file.Name() != "static" {
      // fmt.Println(file.Name())
      gf.NonFatal(os.RemoveAll(file.Name()))
    }
  }
  // gf.NonFatal(os.RemoveAll("bin"))
}
