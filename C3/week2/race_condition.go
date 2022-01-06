package main

import (
	"fmt"
	"time"
)

func printTime(t int, goR int) {
	for i := 0; i < t; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Go routine: %v, time: %v\n", goR, i)
	}
}

func main() {
	fmt.Println("Concurrency")

	// When you execute three goroutine, you will see the order
	// of goR different from times to times.
	// If we take any variable and compute with the printTime function,
	// say take the value of this variable plus the goR,
	// you'll see every time the program is executed
	// the program print out different output.
	// => race condition
	go printTime(5, 1)
	go printTime(6, 2)
	go printTime(7, 3)
	fmt.Scanln()
}
