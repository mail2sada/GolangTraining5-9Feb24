package main

import "fmt"

func main() {
	fmt.Println("Demo: Strings")

	var Str string = "Hello"

	fmt.Println(Str)

	for idx, val := range Str {
		fmt.Printf("%d:%c\n", idx, val)
	}
}
