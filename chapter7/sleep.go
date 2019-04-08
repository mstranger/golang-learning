// p. 218
package main

// TODO: tasks 7.6 - 7.7

import (
	"flag"
	"fmt"
	"time"
)

// go run sleep.go -period 1m
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Waiting %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
