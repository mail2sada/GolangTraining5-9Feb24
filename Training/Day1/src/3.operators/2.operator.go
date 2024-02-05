package main

import "fmt"

func main() {
	fmt.Println("Demo: Relational operators")

	var a, b = 10, 20

	var c = a == b

	fmt.Println(c)

	c = a != b
	fmt.Println(c)

	c = a > b
	fmt.Println(c)

	c = a < b
	fmt.Println(c)

	var d = 30

	d += a
	fmt.Println(d)

}
