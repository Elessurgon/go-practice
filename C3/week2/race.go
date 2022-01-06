package main

import (
	"fmt"
	"time"
)

var i int

func f() {
	i = 4
	fmt.Println("f: ", i)

}

func inc() {
	i++
	fmt.Println("inc: ", i)
}

// EXPLANATION
// Here the race condition happens cause
// when the 2 routines are executed, they share the same
// variable, namely i, upon multiple exections, depending on
// how your OS schedules the threads, one can observer that i either is
// finally assigned 1, 4 or 5

func main() {
	go inc()
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println(i)
}
