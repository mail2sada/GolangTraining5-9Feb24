package main

import "fmt"

func main() {
	fmt.Println("Demo: Nested if")
	// Example variables
	age := 25
	gender := "female"

	// Outer if statement
	if age >= 18 {
		fmt.Println("You are an adult.")

		// Nested if statement
		if gender == "female" {
			fmt.Println("You are a female adult.")
		} else {
			fmt.Println("You are a male adult.")
		}
	} else {
		fmt.Println("You are a minor.")
	}
}
