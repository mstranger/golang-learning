// p. 226
package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

// implement Len,  Less, Swap func for sort.Interface
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	var names = []string{"second", "first", "other"}

	fmt.Println(names)
	// same as below
	// sort.Strings(names)
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}
