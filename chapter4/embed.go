// p. 137
package main

import "fmt"

type Point struct {
  X, Y int
}

type Circle struct {
  // Center Point
  Point
  Radius int
}

type Wheel struct {
  // Circle Circle
  Circle
  Spokes int
}

func main() {
  // var w = Wheel{Circle{Point{8, 8}, 5}, 20}
  var w = Wheel{
    Circle: Circle{
      Point:  Point{X: 8, Y: 8},
      Radius: 5,
    },
    Spokes: 20,
  }

  fmt.Printf("%#v\n", w)
  w.X = 42
  fmt.Printf("%#v\n", w)
}
