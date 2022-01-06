package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func LittleSort(sli []int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i, _ := range sli {
		for j := i; j < len(sli); j++ {
			if sli[j] < sli[i] {
				temp := sli[j]
				sli[j] = sli[i]
				sli[i] = temp
			}
		}
	}
}

func MergeArray(ar1 []int, ar2 []int, result chan []int, wg *sync.WaitGroup) {
	defer wg.Done()

	rs := make([]int, 0)
	i1 := 0
	i2 := 0
	for i := 0; i < len(ar1)+len(ar2); i++ {
		switch {
		case i1 == len(ar1):
			rs = append(rs, ar2[i2])
			i2++
		case i2 == len(ar2):
			rs = append(rs, ar1[i1])
			i1++
		default:
			if ar1[i1] < ar2[i2] {
				rs = append(rs, ar1[i1])
				i1++
			} else {
				rs = append(rs, ar2[i2])
				i2++
			}
		}
	}

	result <- rs
}
func main() {
	var num int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n                      **********************************")
	fmt.Println("		FIRST ENTER THE NUMBER OF NUMBER NEED TO BE SORT!\n                      **********************************")

	fmt.Print("Enter number of numbers need to be sorted: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)

	num, err := strconv.Atoi(text)
	num = int(num)
	if err != nil || num <= 0 {
		fmt.Println("Unvalid Value")
		return
	}

	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	arr3 := make([]int, 0)
	arr4 := make([]int, 0)

	for i := 0; i < num; i++ {
		var temp int
		fmt.Printf(" . The %vth number: ", i+1)
		fmt.Scan(&temp)

		switch i % 4 {
		case 1:
			arr1 = append(arr1, temp)
		case 2:
			arr2 = append(arr2, temp)
		case 3:
			arr3 = append(arr3, temp)
		default:
			arr4 = append(arr4, temp)
		}
	}

	var wg sync.WaitGroup
	wg.Add(4)
	go LittleSort(arr1, &wg)
	go LittleSort(arr2, &wg)
	go LittleSort(arr3, &wg)
	go LittleSort(arr4, &wg)

	wg.Wait()

	ch := make(chan []int, 2)

	wg.Add(2)
	go MergeArray(arr1, arr2, ch, &wg)
	go MergeArray(arr3, arr4, ch, &wg)
	wg.Wait()

	wg.Add(1)
	go MergeArray(<-ch, <-ch, ch, &wg)
	wg.Wait()

	fmt.Println("The array after sorting is: ", <-ch)
}
