package main

// TODO: p. 318, task 9.2

import "fmt"

func main() {
	var x, y int

	go func() {
		x = 1
		fmt.Print("y:", y, " ")
	}()

	go func() {
		y = 1
		fmt.Print("x:", x, " ")
	}()
}
