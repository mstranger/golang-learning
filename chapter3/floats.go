package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Max float32: ", math.MaxFloat32)
	fmt.Println("Max float64: ", math.MaxFloat64)
	fmt.Println()

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eˣ = %8.3f\n", x, math.Exp(float64(x)))
	}
}
