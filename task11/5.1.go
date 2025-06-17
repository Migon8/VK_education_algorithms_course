package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var number int
	fmt.Scan(&number)

	maxProduct := math.MinInt64

	for i := 0; i < n-number+1; i++ {
		product := 1
		for j := i; j < i+number; j++ {
			product *= arr[j]
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	fmt.Println(maxProduct)

}
