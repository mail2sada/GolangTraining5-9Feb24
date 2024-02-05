package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch")

	a := 5

	switch {
	case a < 0:
		fmt.Println("a is negetive")
	case a < 10:
		fmt.Println("a is a natural number")
		fallthrough
	case a < 20:
		fmt.Println("a is number between 11-19")
	default:
		fmt.Println("Un handled")
	}
}
