package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Perim() float64
	Area() float64
}

type Rect struct {
	a, b float64
}

func (r Rect) Perim() float64 {
	return 2*r.a + 2*r.b
}

func (r Rect) Area() float64 {
	return r.a * r.b
}

type Circle struct {
	a float64
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.a
}

func (c Circle) Area() float64 {
	return math.Pi * c.a * c.a
}

func measure(g Geometry) {
	fmt.Println("------------------")
	fmt.Println(g)
	fmt.Println(g.Perim())
	fmt.Println(g.Area())
	fmt.Println("------------------")
}

func main() {
	var t = Rect{3, 4}
	measure(t)

	var c = Circle{4}
	measure(c)
}
