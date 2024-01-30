package main

import "fmt"

func main() {
	fmt.Println("Demo: Airthmatic Operators")

	var a, b int = 100, 200

	sum := a + b
	diff := a - b
	mul := a * b
	div := a / b
	mod := a % b

	fmt.Println("Sum is:", sum, "Diff is:", diff, "Product is:", mul, "Quotient is:", div, "Remainder:", mod)
}
