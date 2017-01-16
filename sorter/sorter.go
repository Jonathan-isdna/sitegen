package sorter

import (
  "sort"
  "github.com/Jonathan-isdna/sitegen/render"
)

type ByDate []render.DPost

func (a ByDate) Len() int { return len(a) }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date.Int() > a[j].Date.Int() }

func Sort(postList []render.DPost) {
  sort.Sort(ByDate(postList))
}
