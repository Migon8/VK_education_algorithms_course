package main

import (
	"fmt"
)

func bubbleSort(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func main() {

	var size int
	fmt.Scan(&size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	bubbleSort(arr)

	for i, num := range arr {
		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Print(num)
	}

	fmt.Println()

}
