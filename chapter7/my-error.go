// custom implementation of the 'errors' package
package main

import "fmt"

type myError interface {
	Error() string
}

type myErrorString struct{ s string }

// NOTE: `&` return address
// and guarantees that New("a") != New("a")
func New(s string) myError {
	return &myErrorString{s}
}

func (e *myErrorString) Error() string {
	return e.s
}

func main() {
	err := New("not found")
	fmt.Printf("Message: %s\n", err)
	fmt.Printf("Type: %T\n", err)
	fmt.Println(New("eof") == New("eof"))
	fmt.Printf("%p\n", New("eof"))
	fmt.Printf("%p\n", New("eof"))
}
