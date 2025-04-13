package main

import (
	"fmt"
)

func main() {

	var size int
	fmt.Scan(&size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	currentSum := 1
	maxSum := 1
	value := arr[0]

	for i := 1; i < size; i++ {
		if arr[i] == arr[i-1] {
			currentSum++

			if currentSum > maxSum {
				maxSum = currentSum
				value = arr[i]
			}
		} else {
			currentSum = 1
		}
	}

	if maxSum == 1 {
		fmt.Print(-1)
	} else {
		fmt.Print(value)
	}

}
