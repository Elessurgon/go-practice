package main

import (
	"fmt"
	"sort"
	"sync"
)

var arr []int
var wg sync.WaitGroup

func sorting(sli []int) {
	sort.Ints(sli)
	wg.Done()
}

func merge(newarr, a1, a2 []int) {
	i, j, k := 0, 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			newarr[k] = a1[i]
			i++
		} else {
			newarr[k] = a2[j]
			j++
		}
		k++
	}
	for i < len(a1) {
		newarr[k] = a1[i]
		i++
		k++
	}
	for j < len(a2) {
		newarr[k] = a2[j]
		j++
		k++
	}
}

func main() {
	var n int
	fmt.Println("Enter the no. of elements")
	fmt.Scan(&n)
	arr = make([]int, n)
	fmt.Println("Enter the elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sub := n / 4
	part1, part2, part3, part4 := arr[:sub], arr[sub:2*sub], arr[2*sub:3*sub], arr[3*sub:n]

	fmt.Println("Subarrays")
	fmt.Printf("%v\n", part1)
	fmt.Printf("%v\n", part2)
	fmt.Printf("%v\n", part3)
	fmt.Printf("%v\n", part4)

	wg.Add(4)
	go sorting(part1)
	go sorting(part2)
	go sorting(part3)
	go sorting(part4)
	wg.Wait()

	fmt.Println("Sorted Subarrays")
	fmt.Printf("%v\n", part1)
	fmt.Printf("%v\n", part2)
	fmt.Printf("%v\n", part3)
	fmt.Printf("%v\n", part4)

	newarr1, newarr2, final := make([]int, len(part1)+len(part2)), make([]int, len(part3)+len(part4)), make([]int, n)
	merge(newarr1, part1, part2)
	merge(newarr2, part3, part4)
	merge(final, newarr1, newarr2)
	fmt.Println("Sorted array")
	for i := range final {
		fmt.Print(final[i], " ")
	}

}
