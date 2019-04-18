// p. 288
package main

// TODO: p. 291, task 8.8

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
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}
