// p. 68, exec 2.2
package conversion

import "fmt"

type Feet float64
type Meters float64

func (f Feet) String() string {
  return fmt.Sprintf("%.5gft", f)
}

func (m Meters) String() string {
  return fmt.Sprintf("%.5gm", m)
}
