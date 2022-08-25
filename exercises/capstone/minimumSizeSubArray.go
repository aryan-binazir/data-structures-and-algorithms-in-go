package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
}

func minSubArrayLen(target int, nums []int) int {
	minContSlice := len(nums) + 1
	anchor, runner := 0, 0
	sum := nums[anchor]

	for runner < len(nums) {
		if sum >= target {
			currentLength := runner - anchor + 1
			if currentLength < minContSlice {
				minContSlice = currentLength
			}
			sum -= nums[anchor]
			anchor++
		} else {
			runner++
			if runner == len(nums) {
				break
			}
			sum += nums[runner]
		}
	}
	if minContSlice > len(nums) {
		minContSlice = 0
	}

	return minContSlice
}
