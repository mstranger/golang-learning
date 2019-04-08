// p. 211
package main

// TODO: tasks 7.1 - 7.3
// TODO: tasks 7.4 - 7.5

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	fmt.Println(c)
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Println(c)
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
