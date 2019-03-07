package main

import (
  "fmt"
  "strconv"
)

func main() {
  x := 234
  y := fmt.Sprintf("%d", x)
  fmt.Println(y, strconv.Itoa(x))
}
