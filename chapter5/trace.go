// p. 180
package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter in %s", msg)
	return func() {
		log.Printf("Out of %s (%s)", msg, time.Since(start))
	}
}
