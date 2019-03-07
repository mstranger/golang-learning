package main

import "fmt"

// some/domain/org.com -> org
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[(i+1):]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	fmt.Println(basename("a.c.d.go"))
	fmt.Println(basename("http://google.com"))
	fmt.Println(basename("abs"))
}
