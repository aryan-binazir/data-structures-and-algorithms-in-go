package main

import "fmt"

func Greet(name string) string {
	var greeting string

	if name == "Johnny" {
		greeting = "Hello, my love!"
	} else {
		greeting = "Hello, " + name + "!"
	}
	return greeting
}

func main() {
	fmt.Println(Greet("Alfred")) // Hello, Alfred!
	fmt.Println(Greet("Johnny")) // Hello, my love!
}
