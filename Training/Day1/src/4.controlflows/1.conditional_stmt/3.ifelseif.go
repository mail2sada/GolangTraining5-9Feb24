package main

import "fmt"

func main() {
	fmt.Println("Demo: if-else-if")

	c, d := 10, 10

	if c > d {
		fmt.Println("c is greater than d")
	} else if c < d {
		fmt.Println("c is less than d")
	} else {
		fmt.Println("Both c and d are equal")
	}
}
