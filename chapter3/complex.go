package main

import "fmt"

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(1i * 1i)

	z := 2 + 3i
	fmt.Println(z)
}
