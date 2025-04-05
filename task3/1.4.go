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

	max := arr[0]

	for i := 0; i < size; i++ {

		if arr[i] > max {
			max = arr[i]
		}

	}

	fmt.Print(max)

}
