// p. 187
package main

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	// defer func() {
	// 	if p := recover(); p != nil {
	// 		fmt.Errorf("inner error: %v", p)
	// 	}
	// }()
	
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x-1)
}
