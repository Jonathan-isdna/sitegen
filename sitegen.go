package main

import (
  "os"
  "net/http"
  "log"
  "github.com/Jonathan-isdna/sitegen/render"
  "github.com/Jonathan-isdna/sitegen/filesearch"
  "github.com/Jonathan-isdna/sitegen/fileio"
  // "github.com/Jonathan-isdna/sitegen/gf"
  "time"
  "fmt"
  // "os"
  // "bytes"
)

func main() {
  tBefore := time.Now()


  // fmt.Println("----- Sitegen -----")
  fileio.ResetBinFolder()
  // // Get list of .md posts
  // postList := filesearch.Search("posts")
  t := render.Init()

  // Render Posts
  postList := filesearch.Search("posts")
  for _, post := range postList {
    fmt.Printf("Rendering: %s\n", post)
    render.Post(t, post)
  }

  // Render Content folder
  contentList := filesearch.Search("content")
  for _, file := range contentList {
    fmt.Printf("Rendering: %s\n", file)
    render.Content(t, file)
  }

  // Copy static to bin...


  tAfter := time.Now()
  fmt.Printf("Finished in %v.\n", tAfter.Sub(tBefore))

  // Serve the site after generating it.
  if len(os.Args) > 1 {
    if os.Args[1] == "serve" {
      fmt.Println("Serving on localhost/:8080")
      fmt.Println("Press ctrl+c to exit")
      log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("bin/"))))
    }
  }
}
