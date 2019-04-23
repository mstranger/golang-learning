// +build ignore
// p. 346
package main

// TODO: p. 352, task 10.4

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
