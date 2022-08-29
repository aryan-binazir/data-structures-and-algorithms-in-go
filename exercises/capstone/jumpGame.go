package main

import "fmt"

func main() {
	arr1 := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(arr1))
	arr2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(arr2))

}

func canJump(nums []int) bool {
	cache := map[int]bool{}
	lastIndex := len(nums) - 1
	currentIndex := 0
	return helper(nums, lastIndex, currentIndex, cache)
}

func helper(nums []int, lastIndex, currentIndex int, cache map[int]bool) bool {

	if val, ok := cache[currentIndex]; ok {
		return val
	}

	if currentIndex == lastIndex {
		return true
	} else if nums[currentIndex] == 0 {
		return false // landing on 0
	} else if currentIndex > lastIndex {
		return false // oob index
	}

	for i := 1; i <= nums[currentIndex]; i++ {
		if helper(nums, lastIndex, currentIndex+i, cache) {
			cache[currentIndex] = true
			return true
		}
	}
	cache[currentIndex] = false
	return false
}
