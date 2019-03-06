package main

import (
	"fmt"
	"./substring"
)

func main() {
	s := "hello, world!"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[:6])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("goodbye" + s[5:])

	// error, string is immutable
	// s[0] = 'H'

	const unformat = `<html>Not formated string's tag</html>`
	fmt.Println(unformat)

	r := "\u4e16\u754c"
	fmt.Println(r)

	var rarr = "\u30d7\u30ed\u30b0\u30e9\u30e0"
	fmt.Println(string(rarr))
	fmt.Println(string(123123))

	fmt.Println(substring.HasPrefix("hello, world", "hell"))
	fmt.Println(substring.HasSuffix("hello, world", "orld"))
	fmt.Println(substring.Contains("hello, world", "ello"))
}
