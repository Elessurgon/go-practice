package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
Having length 3 in the intialization does not make any sense.
I think it should have been capacity=3 and length=0.
But I have used the length as 3 just to respect the instructions.
But then, nothing has been told about initialization of first 3 literals.
Hence, I have used slice literals to initialize first 3 literals.
*/

func main() {
	sli := []int{1, 223, -2}
	var inp string = ""
	for true {
		fmt.Printf("\nEnter the integer to be added into slice:\n")
		fmt.Scan(&inp)
		if strings.Compare(inp, "X") != 0 {

			val, err := strconv.Atoi(inp)
			if err != nil {
				fmt.Printf("Wrong Input! Please Check\n")
				continue
			}
			sli = append(sli, val)

			fmt.Printf("Before sorting the slice:\n")
			for _, v := range sli {
				fmt.Printf("%d ", v)
			}
			fmt.Printf("\n")
			sort.Slice(sli, func(i, j int) bool {
				return sli[i] < sli[j]
			})
			fmt.Printf("Sorted order of Slice:\n")
			for _, v := range sli {
				fmt.Printf("%d ", v)
			}
			fmt.Printf("\n")
		} else {
			break
		}
	}

}
