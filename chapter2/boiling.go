// p. 51
package main

import "fmt"

const boilingF = 212.0

// NOTE:
//   %g - floating-point, %e for large, %f otherwise
//   %e - scientific notation 1.234e+79
//   %f - decimal without exponent 123.45

func main() {
  var f = boilingF
  var c = (f - 32) * 5 / 9
  fmt.Printf("Температура кипения = %g°F или %g°C\n", f, c)
}
