package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch with relational operator")
	num := 7

	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num >= 0 && num < 10:
		fmt.Println("Single-digit positive number")
	case num >= 10 && num < 100:
		fmt.Println("Double-digit positive number")
	default:
		fmt.Println("Number is 100 or greater")
	}
}
