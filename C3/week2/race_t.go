package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var x int64 = 1
	for i := 0; i < 3; i++ {
		go double(&x)

		//time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		/**
		removing the above line causes race condition, i.e.
		double and square func both run in unpredictable sequence
		rendering different final value of x, each time the program is
		run
		To check the race condition, run
		go run -race Race.go
		**/

		go square(&x)

	}
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	fmt.Printf("%s %d", "\nThe final value of x is: ", x)

}

// Doubles the value of x
func double(x *int64) {

	*x = *x * 2
	fmt.Printf("%s %d", "\nx is doubled to: ", *x)
}

// Squares the value of x
func square(x *int64) {
	*x = *x * *x
	fmt.Printf("%s %d", "\nx is squared to: ", *x)

}
