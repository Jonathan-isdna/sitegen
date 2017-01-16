package render

import (
  "html/template"
  "bytes"
  "github.com/russross/blackfriday"
  "github.com/Jonathan-isdna/sitegen/fileio"
  "github.com/Jonathan-isdna/sitegen/gf"
  "github.com/Jonathan-isdna/sitegen/filesearch"
  "github.com/Jonathan-isdna/sitegen/datetime"
  "path/filepath"
  // "fmt"
)

type DPage struct {
  Title, Url, HtmlFile string
}

type DContent struct {
  postList []DPost
}

type DPost struct {
  DPage
  Date datetime.Date
  Image string
  Post template.HTML
  MDFile string
  ID int
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

func ParsePost(postMd string) (pData DPost) {
  var postText string
  lines := fileio.FileReadLines(postMd)
  if len(lines) < 4 { gf.Generate("Markdown error. Last file can't be parsed") }

  pData.Title = lines[0]
  pData.Date = datetime.ParseDate(lines[1])
  pData.Image = lines[2]
  pData.MDFile = postMd
  pData.Url = "blog/" + postMd[:len(postMd)-3]
  pData.HtmlFile = "bin/" + pData.Url + "/index.html"
  for i := 3; i < len(lines); i++ { postText += lines[i] + "\r\n" }
  pData.Post = template.HTML(blackfriday.MarkdownBasic([]byte(postText)))
  return pData
}

func Content(t *template.Template, file string, data interface{}) {
  _, tpl := filepath.Split(file)
  filename := file[8:len(file)]
  var compiledFilename string
  if tpl != "index.html" {
    compiledFilename = "bin/" + filename[:len(filename)-5] + "/index.html"
  } else {
    compiledFilename = "bin/index.html"
  }
  fileio.FileWrite(compiledFilename, GetHtml(t, tpl, data))
}

func Init() *template.Template {
  // var t *template.Template
  content := filesearch.Search("content")
  templates := filesearch.Search("templates")

  paths := append(templates, content...)
  t := template.Must(template.ParseFiles(paths...))
  return t
}

func GetHtml(t *template.Template, tpl string, data interface{}) string {
  var html bytes.Buffer
  err := t.ExecuteTemplate(&html, tpl, data)
  gf.Check(err)
  return html.String()
}
