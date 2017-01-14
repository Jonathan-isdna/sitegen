package render

import (
  "html/template"
  "bytes"
  "github.com/russross/blackfriday"
  "github.com/Jonathan-isdna/sitegen/fileio"
  "github.com/Jonathan-isdna/sitegen/gf"
  "github.com/Jonathan-isdna/sitegen/filesearch"
  "path/filepath"
  // "fmt"
)

type dPost struct {
  Title string
  Date string
  Image string
  Post template.HTML
}

// Takes a string template input and data and return html string back
func renderTemplate(tpl string, data interface{}) string {
  // Create template object
  t := template.New("post")

  // Parse template
  t, err := t.Parse(tpl)
  gf.Check(err)

  // Render template
  var rendered bytes.Buffer
  t.Execute(&rendered, data)

  // Return template as string
  return rendered.String()
}

func parseMarkdown(lines []string) (Title string, Date string, Image string, Post template.HTML) {
  if len(lines) < 4 {
    gf.Generate("Markdown error. Last file can't be parsed")
  }
  Title = lines[0]
  Date = lines[1]
  Image = lines[2]
  postText := ""
  for i := 3; i < len(lines); i++ { postText += lines[i] + "\r\n" }
  Post = template.HTML(blackfriday.MarkdownBasic([]byte(postText)))
  return Title, Date, Image, Post
}

// func Post(postFile string) {
//   var data dPost
//   if postFile[len(postFile)-3:] != ".md" {
//     gf.Generate("\nError processing markdown files.\nAre they named correctly?")
//   }
//   postName := postFile[:len(postFile)-3]
//
//   postMdLines := fileio.FileReadLines(postFile)
//   data.Title, data.Date, data.Image, data.Post = parseMarkdown(postMdLines)
//
//   postTemplate := fileio.FileRead("templates/post.html")
//   rendered := renderTemplate(postTemplate, data)
//   filename := "bin/blog/" + postName + "/index.html"
//   fileio.FileWrite(filename, rendered)
// }

// var indexPage bytes.Buffer
// err := templates.ExecuteTemplate(&indexPage, "index.html", nil)
// gf.Check(err)
// fileio.FileWrite("bin/index.html", indexPage.String())

func Post(t *template.Template, postFile string) {

  var data dPost
  if postFile[len(postFile)-3:] != ".md" {
    gf.Generate("\nError processing markdown files.\nAre they named correctly?")
  }
  postName := postFile[:len(postFile)-3]
  postMdLines := fileio.FileReadLines(postFile)
  data.Title, data.Date, data.Image, data.Post = parseMarkdown(postMdLines)

  filename := "bin/blog/" + postName + "/index.html"
  fileio.FileWrite(filename, getHtml(t, "post", data))
}

func Content(t *template.Template, file string) {
  _, tpl := filepath.Split(file)
  filename := file[8:len(file)]
  var compiledFilename string
  if tpl != "index.html" {
    compiledFilename = "bin/" + filename[:len(filename)-5] + "/index.html"
  } else {
    compiledFilename = "bin/index.html"
  }
  fileio.FileWrite(compiledFilename, getHtml(t, tpl, nil))
}

func Init() *template.Template {
  // var t *template.Template
  content := filesearch.Search("content")
  templates := filesearch.Search("templates")

  paths := append(templates, content...)
  t := template.Must(template.ParseFiles(paths...))
  return t
}

func getHtml(t *template.Template, tpl string, data interface{}) string {
  var html bytes.Buffer
  err := t.ExecuteTemplate(&html, tpl, data)
  gf.Check(err)
  return html.String()
}
