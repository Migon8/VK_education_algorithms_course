package main

import (
	"fmt"
	"sort"
)

func FindExponentialBounds(arr []int, target int) (int, int) {
	if len(arr) == 0 || target < arr[0] {
		return -1, -1
	}

	i := 1
	for i < len(arr) && arr[i] <= target {
		i *= 2
	}
	left := i / 2
	right := min(i, len(arr)-1)

	return left, right

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

	if !sort.IntsAreSorted(arr) {
		fmt.Println("Массив должен быть отсортирован!")
		return
	}

	left, right := FindExponentialBounds(arr, target)

	fmt.Print(left, right)

}
