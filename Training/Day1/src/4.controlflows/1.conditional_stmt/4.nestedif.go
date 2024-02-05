package main

import "fmt"

func main() {
	fmt.Println("Demo:Nested if")
	a, b, c := 10, 20, 30

	if a > b {
		if a > c {
			fmt.Println("a is biggest")
		} else if c > a {
			fmt.Println("c is biggest")
		} else {
			fmt.Println("both c and a are equal")
		}
	} else if b > c {
		fmt.Println("b is biggest")
	} else if c > b {
		fmt.Println("c is biggest")
	} else {
		fmt.Println("Both c and b are equal")
	}
}
