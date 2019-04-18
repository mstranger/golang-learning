// p. 288
package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Start")
}

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read 1 byte
		abort <- struct{}{}
	}()

	fmt.Println("Begin countdown. Press <enter> for abort...")

	select {
	case <-time.After(10 * time.Second):
		// do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
