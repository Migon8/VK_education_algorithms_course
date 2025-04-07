package main

import (
	"fmt"
)

func binarySearch(arr []int, index int) (int, bool) {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == index {
			return mid, true
		} else if arr[mid] < index {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left, false

}

func main() {

	var size int
	fmt.Scan(&size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	var target int
	fmt.Scan(&target)

	pos, _ := binarySearch(arr, target)

	fmt.Print(pos)

}
