package main

import "fmt"

/*
	for <Intialization>;<Condition>;Increment {
		// Statements
	}
*/

func main() {
	fmt.Println("Demo: For loop")

	for i := 0; i < 10; i++ {
		fmt.Println("Value of i is", i)
	}

	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}

	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
