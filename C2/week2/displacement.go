package main

import (
	"fmt"
)

func GenDisplaceFn(acc float64, init_vel float64, init_disp float64) func(float64) float64 {
	return func(tm float64) float64 {
		return 0.5*acc*tm*tm + init_vel*tm + init_disp
	}
}

func main() {
	// acceleration, initial velocity, and initial displacement.
	var acc, init_vel, init_disp float64
	fmt.Println("Enter accleration")
	fmt.Scanln(&acc)
	fmt.Println("Enter initial velocity")
	fmt.Scanln(&init_vel)
	fmt.Println("Enter initial displacement")
	fmt.Scanln(&init_disp)
	fn := GenDisplaceFn(acc, init_vel, init_disp)
	for true {
		var tm float64
		fmt.Println("Enter time or Ctrl + C to break")
		fmt.Scanln(&tm)
		fmt.Println("Total displacement: ", fn(tm))
	}
}
