package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) //[[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))             //[]
	fmt.Println(threeSum([]int{0, 0, 0}))             //[[0,0,0]]
}

func threeSum(slice []int) [][]int {
	result := [][]int{}
	sliceLength := len(slice)

	for i := 0; i <= sliceLength-3; i++ {
		for j := i + 1; j <= sliceLength-2; j++ {
			for k := j + 1; k <= sliceLength-1; k++ {

				if slice[i]+slice[j]+slice[k] == 0 {
					addSlice := true
					potentialSlice := []int{slice[i], slice[j], slice[k]}
					for _, subSl := range result {
						if sameValues(subSl, potentialSlice) {
							addSlice = false
						}
					}

					if addSlice {
						result = append(result, potentialSlice)
					}
				}

			}
		}
	}
	return result
}

func sameValues(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Ints(a)
	sort.Ints(b)

	for index, value := range a {
		if !(value == b[index]) {
			return false
		}
	}

	return true
}
