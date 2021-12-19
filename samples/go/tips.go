package main

func main() {
}

// sort
import "sort"

type Sortable []string

func (s *Sortable) Len() int {}
func (s *Sortable) Swap(i, j) {s[i], s[j] = s[j], s[i]}
func (s *Sortable) Less(i, j) bool {} // returns true if less

// also available
// sort.Ints(a)  -- an array is already sorted after this
// sort.Float64s
// sort.Strings

// also available
sort.SliceStable(array, func /*Less*/(i, j int) bool {return a[i] < a[j]})


