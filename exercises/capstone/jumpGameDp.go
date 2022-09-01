package main

import (
	"fmt"
)

func canJump(nums []int) bool {

	cache := map[int]bool{}
	cache[len(nums)-1] = true // base case
	for pointer := len(nums) - 1; pointer >= 0; pointer-- {
		if nums[pointer] == 0 && pointer != len(nums)-1 {
			cache[pointer] = false
			continue
		}
		for toAdd := 1; toAdd <= nums[pointer]; toAdd++ {
			if pointer+toAdd > len(nums) {
				break
			}
			if cache[pointer+toAdd] == true {
				cache[pointer] = true
				break
			}
		}
	}
	return cache[0]
}

func main() {
	arr1 := []int{2, 3, 1, 1, 4} // true
	fmt.Println(canJump(arr1))

	arr2 := []int{3, 2, 1, 0, 4} // false
	fmt.Println(canJump(arr2))

	arr3 := []int{2, 0}        // true
	fmt.Println(canJump(arr3)) // true

	arr4 := []int{0, 1, 2, 3} // false
	fmt.Println(canJump(arr4))

	arr5 := []int{4, 0, 0, 0, 1} // true
	fmt.Println(canJump(arr5))
}
