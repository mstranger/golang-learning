package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // position or -1
	s = s[(slash + 1):]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("a.c.d.go"))
	fmt.Println(basename("http://google.com"))
	fmt.Println(basename("abs"))	
}
