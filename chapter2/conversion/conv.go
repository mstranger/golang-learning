// p. 68, exec. 2.2
package conversion

const (
  FeetInMeter Meters = 0.3048
)

// convert feet to meters
func FToM(ft Feet) Meters {
  return Meters(ft * Feet(FeetInMeter))
}

// convert meters to feet
func MToF(m Meters) Feet {
  return Feet(m / FeetInMeter)
}
