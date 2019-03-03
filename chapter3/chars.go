package main

import "fmt"

func main() {
	for i := 97; i <= 122; i++ {
		fmt.Printf("%d-%[1]c-%#[1]x\n", i)
	}

	fmt.Println()
}
