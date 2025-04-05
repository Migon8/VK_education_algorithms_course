package main

import "fmt"

func changeZero(arr []int) {
	nonZero := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[nonZero], arr[i] = arr[i], arr[nonZero]
			nonZero++
		}
	}
}

func main() {
	var size int
	fmt.Scan(&size)

	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	changeZero(numbers)

	for i, num := range numbers {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}
