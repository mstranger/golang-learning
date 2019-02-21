// p. 56
package main

import (
  "flag"
  "fmt"
  "strings"
)

var n = flag.Bool("n", false, "пропуск симвла новой строки")
var sep = flag.String("s", " ", "разделитель")

// go build
// echo4 -help [-s, -n]
// echo4 -s / a bc def

func main() {
  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }
}
