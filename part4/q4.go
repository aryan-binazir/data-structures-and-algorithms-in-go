package main

import "fmt"

func main() {
	fmt.Println(greatestNumber([]int{1, 2, 12, 4, 5, 6})) // 12
}

// refactoring to O(N)
func greatestNumber(array []int) int {
	greatestIValIndex := array[0]

	for _, value := range array {
		if value > greatestIValIndex {
			greatestIValIndex = value
		}
	}

	return greatestIValIndex
}
