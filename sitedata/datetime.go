package datetime

import (
  "strings"
  "github.com/Jonathan-isdna/sitegen/gf"
)

type Date struct {
  Year int
  Month int
  Day int
}

func ParseDate(dateString string) (pd date) {
  var ds [3]int
  var err error
  for i, v := range strings.Split(dateString, "/") {
    ds[i], err = strconv.Atoi(v)
    check(err)
  }
  pd.Month, pd.Day, pd.Year = ds[0], ds[2], ds[3]
  return pd
}
