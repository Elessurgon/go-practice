package main

import (
	"fmt"
	s "strings"
)

func main() {
	var input string
	fmt.Printf("Enter input string: ")
	fmt.Scanln(&input)
	if s.HasPrefix(input, "i") && s.HasSuffix(input, "n") && s.Contains(input, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
