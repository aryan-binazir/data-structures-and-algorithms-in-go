package main

import "fmt"

func subsets(nums []int) [][]int {
	var results [][]int
	var candidate []int
	backtrack(nums, candidate, &results)
	return results
}

func backtrack(list []int, candidate []int, results *[][]int) {
	if allUnique(candidate) { // also can return if terminal condition (lead node)
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
		// return
	}

	for i := 0; i < len(list); i++ {
		// if { //<<don't want to take; redundant to dead-end condition above (only need 1)
		//   continue
		// }
		if len(list) == 0 {
			return
		}
		candidate = append(candidate, list[i])    // take
		backtrack(list[i+1:], candidate, results) // explore
		candidate = candidate[:len(candidate)-1]  // clean up
	}
}

func main() {
	arr1 := []int{1, 2, 3}

	fmt.Println(subsets(arr1))
}

func allUnique(arr []int) bool {
	hsh := map[int]bool{}

	for _, v := range arr {
		_, ok := hsh[v]
		if !ok {
			hsh[v] = true
		} else {
			return false
		}
	}
	return true
}
