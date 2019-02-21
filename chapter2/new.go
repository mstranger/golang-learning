// p. 57
package main

import "fmt"

func main() {
  p := new(int)
  fmt.Println(p)
  fmt.Println(*p)
  *p = 2
  fmt.Println(*p)

  q := new(int)
  fmt.Println(q)
  fmt.Println(q == p)
}
