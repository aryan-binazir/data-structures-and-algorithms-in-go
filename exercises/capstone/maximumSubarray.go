package main

import "fmt"

func main() {

}

func maxSubArray(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	midPoint := len(arr) / 2
	leftMax := maxSubArray(arr[:midPoint])
	rightMax := maxSubArray(arr[midPoint:])
	crossMax := maxCross(arr)
	fmt.Println(leftMax, rightMax, crossMax)
	return maxVal([]int{leftMax, rightMax, crossMax})
}

func maxCross(arr []int) int {
	runningLeftSum := 0
	maxLeftSum := -999999
	mid := len(arr) / 2

	for i := mid - 1; i >= 0; i-- {
		runningLeftSum += arr[i]

		if runningLeftSum > maxLeftSum {
			maxLeftSum = runningLeftSum
		}
	}

	runningRightSum := 0
	maxRightSum := -999999

	for i := mid; i <= len(arr)-1; i++ {
		runningRightSum += arr[i]

		if runningRightSum > maxRightSum {
			maxRightSum = runningRightSum
		}
	}
	return maxLeftSum + maxRightSum
}

func maxVal(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
