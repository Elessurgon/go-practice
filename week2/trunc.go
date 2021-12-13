package main

import "fmt"

func main() {
	fmt.Println("Enter a floating point number")
	var num float64
	fmt.Scanln(&num)
	var num_trunc int64 = int64(num)
	fmt.Printf("Truncated number %v", num_trunc)
}
