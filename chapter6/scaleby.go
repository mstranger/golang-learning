// p. 194
package main

// NOTE: again pointers, p 195

import (
	"fmt"

	"./geometry"
)

func main() {
	p := geometry.Point{1, 2}
	pptr := &p

	fmt.Println(p)
	pptr.ScaleBy(2)
	// imlicit (&p).ScaleBy
	p.ScaleBy(2)
	fmt.Println(p)

	// allow
	q := geometry.Point{3, 4}.Distance(geometry.Point{5, 6})
	fmt.Println(q)

	// not allow (!), temporary without address
	// q = geometry.Point{3, 4}.ScaleBy(2)
}
