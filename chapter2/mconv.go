// p. 68, exec. 2.2
package main

import (
  "fmt"
  "os"
  "strconv"

  "./conversion"
)

// TODO: add cm, d, other
func main() {
  a := os.Args[1]
  b := os.Args[2]
  for _, arg := range os.Args[3:] {
    t, err := strconv.ParseFloat(arg, 64)
    if err != nil {
      fmt.Fprintf(os.Stderr, "mconv: %v\n", err)
      os.Exit(1)
    }

    if a == "m" && b == "ft" {
      m := conversion.Meters(t)
      ft := conversion.MToF(m)
      fmt.Printf("%s\t->\t%s\n", m, ft)
    } else if a == "ft" && b == "m" {
      ft := conversion.Feet(t)
      m := conversion.FToM(ft)
      fmt.Printf("%s\t->\t%s\n", ft, m)
    } else {
      fmt.Println("use format <a b number>, where a, b in [m, ft]")
    }

  }
}
