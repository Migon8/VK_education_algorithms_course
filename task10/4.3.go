package main

import "fmt"

func main() {

	var m string

	fmt.Scan(&m)

	counts := make(map[rune]int)

	for _, ch := range m {
		counts[ch]++
	}

	maxCount := 0

	for _, count := range counts {
		if count > maxCount {
			maxCount = count

		}
	}

	fmt.Print(maxCount)

}
