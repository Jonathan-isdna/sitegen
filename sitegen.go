package main

import (
  "os"
  "net/http"
  "log"
  "github.com/Jonathan-isdna/sitegen/render"
  "github.com/Jonathan-isdna/sitegen/filesearch"
  "github.com/Jonathan-isdna/sitegen/fileio"
  "github.com/Jonathan-isdna/sitegen/gf"
  "github.com/Jonathan-isdna/sitegen/sorter"
  "time"
  "fmt"
  // "os"
  // "bytes"
)

func main() {
  tBefore := time.Now()


  fmt.Println("----- Sitegen -----")
  fileio.ResetBinFolder()
  t := render.Init()

  // Get List of posts and a list of content
  postFileList := filesearch.Search("posts")
  contentList := filesearch.Search("content")

  // String Len for printing current post stuff?
  slen := 10

  // Render Posts
  var data render.DPost
  var postList []render.DPost
  for index, post := range postFileList {
    if len(post) < slen { fmt.Printf("\rPost: %s", post) } else { fmt.Printf("\rPost: %s", post[:slen]) }

    if post[len(post)-3:] != ".md" { gf.Generate("\nError processing markdown files.\nAre they named correctly?") }
    data = render.ParsePost(post)
    data.ID = index
    postList = append(postList, data)
    fileio.FileWrite(data.HtmlFile, render.GetHtml(t, "post", data))
  }
  // Sort postList
  sorter.Sort(postList)
  // for _, post := range postList {
  //   fmt.Printf("Post %d | %d\n", post.ID, post.Date.Int())
  // }

  // Render Content folder
  for _, file := range contentList {
    if len(file) < slen { fmt.Printf("\rCont: %s", file) } else { fmt.Printf("\rContent: %s", file[:slen]) }
    // fmt.Printf("\rContent: %s", file)
    render.Content(t, file, postList)
  }

  // Copy static to bin

  tAfter := time.Now()
  fmt.Printf("\r----- Finished in %v. -----\n", tAfter.Sub(tBefore))

  // Serve the site after generating it.
  if len(os.Args) > 1 {
    if os.Args[1] == "serve" {
      fmt.Println("Serving on localhost/:8080")
      fmt.Println("Press ctrl+c to exit")
      log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("bin/"))))
    }
  }
}
