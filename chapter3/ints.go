package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	fmt.Println("----------")
	fmt.Printf("%v in binary: %08b\n", 21, 21)

	f := 2.99
	
	fmt.Println(int(f))

	fmt.Printf("%o %#[1]o\n", 24)

	ch := 'r'
	fmt.Printf("%d %[1]c %[1]q\n", ch)

	fmt.Printf("%c\n", 9733)
}
