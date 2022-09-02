package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var results [][]int // stores subarr of sum values
	var potential []int
	backtrack(candidates, potential, &results, target)
	return results

}

func backtrack(list []int, potential []int, results *[][]int, target int) {
	if calcSum(potential) == target { // also can return if terminal condition (lead node) // if end of path, (checks whether sum >= target)
		copyPotential := make([]int, len(potential))
		copy(copyPotential, potential)
		*results = append(*results, copyPotential)
		return
	}

	for i := 0; i < len(list); i++ {
		if calcSum(potential)+list[i] > target {
			continue
		}

		// if sum + currNum > target
		// sum += num

		potential = append(potential, list[i])          // take
		backtrack(list[i:], potential, results, target) // explore
		potential = potential[:len(potential)-1]        // clean up
	}
}

func calcSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	arr1 := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(arr1, target))
}
