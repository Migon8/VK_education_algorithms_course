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

	var lastmax int
	found := false

	var number int = -1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i]%2 == 0 {
			lastmax = arr[i]
			found = true
			break
		}
	}

	if found {
		fmt.Print(lastmax)
	} else {

		fmt.Print(number)

	}

}
