// p. 271
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// generate ints
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// power
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// print in the main program
	for {
		fmt.Println(<-squares)
	}
}
