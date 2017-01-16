package datetime

import (
  "strings"
  "strconv"
  "github.com/Jonathan-isdna/sitegen/gf"
)

type Date struct {
  Year, Month, Day int
}

func ParseDate(dateString string) (d Date) {
  var ds [3]int
  var err error

  for i, v := range strings.Split(dateString, "/") {
    // Atoi
    ds[i], err = strconv.Atoi(v)
    gf.Check(err)
  }
  d.Month, d.Day, d.Year = ds[0], ds[1], ds[2]
  return d
}

func (d Date) Int() (i int) { return (d.Year * 10000) + (d.Month * 100) + d.Day }
func (d Date) Show() (s string) {
  s = strconv.Itoa(d.Month) + strconv.Itoa(d.Day) + strconv.Itoa(d.Year)
  return s
}
