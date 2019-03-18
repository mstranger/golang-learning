// p. 120
// TODO: task 4.3 - 4.7

package main

import "fmt"

func nonemty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// with 'append' function
func nonemty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonemty2(data))
	fmt.Printf("%q\n", data)
}

