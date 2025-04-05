package main

import (
	"fmt"
)

func deleteNumber(slice []int, value int) []int {
	i := 0
	for _, v := range slice {
		if v != value {
			slice[i] = v
			i++
		}
	}
	return slice[:i]
}

func main() {

	var size int
	fmt.Scan(&size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	var number int
	fmt.Scan(&number)

	arr = deleteNumber(arr, number)

	for i, num := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}
