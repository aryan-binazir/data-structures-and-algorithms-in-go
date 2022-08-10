package main

import "fmt"

func main() {
	fmt.Println(sum(1, 10)) // 55
}

func sum(low int, high int) int {
	if high == low {
		return low
	}
	return high + sum(low, high-1)
}
