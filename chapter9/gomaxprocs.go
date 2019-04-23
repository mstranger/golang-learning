package main

// TODO: p. 330, task 9.6

import "fmt"

// NOTE: GOMAXPROCS=2 go run ...

func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
