package main

import "fmt"

func main() {
	fmt.Println("Demo operators")

	var a, b = 200, 100

	var c int

	c = a + b
	fmt.Println("Sum is ", c)

	c = a - b
	fmt.Println("Difference", c)

	c = a * b
	fmt.Println("Product is", c)
	c = a / b
	fmt.Println("Qotient", c)

	c = a % b
	fmt.Println("Remainder", c)
}
