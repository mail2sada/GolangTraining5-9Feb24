package main

import "fmt"

const PI = 3.14159

func Circle(radius float64) (float64, float64) {
	cir := 2 * PI * radius

	area := PI * radius * radius

	return cir, area
}

func main() {
	fmt.Println("Demo: Multiple returns")

	cir, area := Circle(2.0)

	fmt.Println(cir, area)

	_, area = Circle(1.0)

	fmt.Println(area)

	cir, _ = Circle(3.0)
	fmt.Println(cir)
}
