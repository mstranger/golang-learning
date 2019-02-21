// p. 65
package tempconv

// CToF convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
  return Fahrenheit(c * 9 / 5 + 32)
}

// FToC convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
  return Celsius((f - 32) * 5 / 9)
}

// exec. 2.1
// CToK convert Celsius to Kelvin
func CToK(c Celsius) Kelvin {
  return Kelvin(c + 273.15)
}

// KToC convert Kelvin to Celsius
func KToC(k Kelvin) Celsius {
  return Celsius(k - 273.15)
}
