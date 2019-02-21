// run 'tempconv' package
package main

import (
  "./tempconv"
  "fmt"
)

func main() {
  fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
  fmt.Println(tempconv.CToF(tempconv.BoilingC))
  fmt.Printf("Boiling in Kelvin scale: %v\n", tempconv.CToK(tempconv.BoilingC))
}
