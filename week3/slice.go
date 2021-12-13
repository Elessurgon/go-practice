package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	a := make([]int, 0, 3)
	for {
		fmt.Println("Enter a number or X to Quit")
		var s string
		fmt.Scanln(&s)
		if s == "X" || s == "x" {
			break
		}
		i, _ := strconv.Atoi(s)
		a = append(a, i)
		sort.Ints(a)
		fmt.Println(a)
	}
}
