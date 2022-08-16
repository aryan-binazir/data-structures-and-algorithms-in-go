package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(comp(a1, a2)) // true

	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}

	fmt.Println(comp(a, b)) // true
}

func comp(arr1 []int, arr2 []int) bool {
	sort.Ints(arr1)
	sort.Ints(arr2)

	for index := range arr1 {
		if int(math.Pow(float64(arr1[index]), 2)) != arr2[index] {
			return false
		}
	}
	return true
}
