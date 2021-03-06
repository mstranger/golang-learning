package main

import "fmt"

// p. 99
// comma use recursion
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("12312321"))
	fmt.Println([]byte("abc"))
	fmt.Println(string([]byte("abc")))
}
