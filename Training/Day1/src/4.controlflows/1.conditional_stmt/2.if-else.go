package main

import "fmt"

func main() {
	fmt.Println("Demo if else")

	/*
		if condition {
			condition is true
		} else {
			condition is false
		}
	*/

	a, b := 10, 20

	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("a is less than b")
	}

	c, d := 10, 10

	if c > d {
		fmt.Println("c is greater than d ")
	} else {
		fmt.Println("c is less than d")
	}
}
