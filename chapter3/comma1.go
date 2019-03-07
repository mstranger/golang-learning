package main

import (
	"bytes"
	"fmt"
)

// comma1 use itteration and buffer type
// p. 101, task 3.10
func comma1(s string) string {
	n := len(s)
	var cnt int

	var buf bytes.Buffer

	for i := n - 1; i >= 0; i-- {
		cnt++
		if cnt == 3 {
			buf.WriteString(",")
			cnt = 0
		}
		fmt.Fprintf(&buf, "%c", s[n-i-1])
	}

	return buf.String()
}

func main() {
	fmt.Println(comma1("12345"))
}
