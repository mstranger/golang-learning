// p. 271
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// generate ints
	go func() {
		// defer close(naturals)
		for x := 0; x < 20; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// power
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// print in the main program
	for x := range squares {
		fmt.Println(x)
	}
}
