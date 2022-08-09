package main

import "fmt"

func main() {
	fmt.Println(findDupVal([]string{"a", "b", "c", "d", "b", "e", "f"})) //b
	fmt.Println(findDupVal([]string{"a", "b", "c", "d", "d", "e", "f"})) //d
}

func findDupVal(strArr []string) string {
	repeatedStr := ""
	stringsFound := make(map[string]bool)

	for _, currStr := range strArr {
		if stringsFound[currStr] {
			repeatedStr = currStr
			break
		}
		stringsFound[currStr] = true
	}
	return repeatedStr
}
