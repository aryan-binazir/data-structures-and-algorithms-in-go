package main

import "fmt"

func main() {
	// fmt.Println(findPeakElement([]int{1, 2, 3, 1}))          //2
	// fmt.Println(findPeakElement([]int{1, 2, 3, 4, 7, 2, 1})) //4
	fmt.Println(findPeakElement([]int{1, 2})) //1

}

func findPeakElement(nums []int) int {
	var index int

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)
		fmt.Println("mid", mid)
		prev := mid - 1
		next := mid + 1

		fmt.Println(left, right)

		if (mid == 0 || mid == len(nums)-1) && !(len(nums) > 2) {
			return mid
		} else if nums[mid] > nums[prev] && nums[mid] > nums[next] {
			index = mid
			break
		} else if nums[mid] > nums[prev] && nums[mid] < nums[next] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index

}

/*
- Input nums array
- choose midpoint
- check number and before and after
	- if both are lower return index of current position
	- if lower index is lower go left
	- if upper index is higher go right
	-

*/

// package main Aashish

// import "fmt"

// func main() {
// 	fmt.Println(findPeakElement([]int{1, 2, 3, 1})) //2
// 	fmt.Println(findPeakElement([]int{4, 3, 2, 1})) //2
// }
// func findPeakElement(nums []int) int {
// 	left, right := 0, len(nums)-1
// 	for left < right {
// 		mid := left + ((right - left) / 2)
// 		if nums[mid] < nums[mid+1] {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}
// 	}
// 	return left
// }
