package main

import "fmt"

func main() {

	fmt.Println("Demo: if-else if-else")

	x := 100
	if x > 100 {
		fmt.Println("X is greater than 100")
	} else if x < 100 {
		fmt.Println("X is less than 100")
	} else {
		fmt.Println("X is equal to 100")
	}

}
