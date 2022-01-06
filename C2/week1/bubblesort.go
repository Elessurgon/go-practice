package main

import (
	"fmt"
)

var n int

func Swap(sli []int, index int) {
	fmt.Println(index)
	temp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = temp
}

func BubbleSort(sli []int) {
	size := len(sli)
	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {

	fmt.Println("Enter number of elements")
	fmt.Scanln(&n)
	fmt.Println("Enter the elements")
	sli := make([]int, n, 10)
	for i := 0; i < n; i++ {
		fmt.Scanln(&sli[i])
	}
	BubbleSort(sli)
	fmt.Println(sli)

}
