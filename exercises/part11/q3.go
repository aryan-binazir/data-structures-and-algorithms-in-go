package main

import "fmt"

func main() {
	fmt.Println(triangleNum(7)) // 28
	fmt.Println(triangleNum(3)) // 6
	fmt.Println(triangleNum(4)) // 10
}

func triangleNum(n int) int {
	if n == 1 {
		return 1
	}

	return n + triangleNum(n-1)
}
