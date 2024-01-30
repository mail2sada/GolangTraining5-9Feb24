package main

import "fmt"

func main() {
	fmt.Println("Demo: Relational Operators")
	x := 5
	y := 8
	isEqual := x == y          // false
	isNotEqual := x != y       // true
	isLess := x < y            // true
	isGreater := x > y         // false
	isLessOrEqual := x <= y    // true
	isGreaterOrEqual := x >= y // false
	fmt.Println(isEqual, isNotEqual, isLess, isGreater, isLessOrEqual, isGreaterOrEqual)
}
