// p. 202
package main

import (
	"fmt"

	"./geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	distFromP := p.Distance
	fmt.Println(distFromP)
	fmt.Println(distFromP(q))
	var zero geometry.Point
	fmt.Println(distFromP(zero))

	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)
	fmt.Println(p)
}
