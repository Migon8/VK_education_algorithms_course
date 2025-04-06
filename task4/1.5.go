package main

import (
	"fmt"
	"math"
)

func main() {
	var size int
	fmt.Scan(&size)
	if size < 2 {
		return
	}

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	minDiff := math.MaxInt64
	var firstPair [2]int

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			minDiff = diff
			firstPair = [2]int{arr[i], arr[i+1]}

		}
	}

	fmt.Println(firstPair[0], firstPair[1])
}
