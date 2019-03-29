// p. 169
// anonym func and closures example
package main

import "fmt"

func squrares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squrares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
