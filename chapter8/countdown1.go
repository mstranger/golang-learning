// p. 288
package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Start")
}

func main() {
	fmt.Println("Begin countdown...")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}
