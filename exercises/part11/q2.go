package main

import "fmt"

func main() {
	fmt.Println(getEvenNum([]int{1, 2, 3, 4, 5, 6, 7})) // [2,4,6]
	fmt.Println(getEvenNum([]int{1, 3, 5}))             // []
	fmt.Println(getEvenNum([]int{}))                    // []
}

func getEvenNum(numSlice []int) []int {
	if len(numSlice) == 0 {
		return []int{}
	}

	if numSlice[0]%2 == 0 {
		return append(getEvenNum(numSlice[1:]), numSlice[0])
	} else {
		return getEvenNum(numSlice[1:])
	}
}
