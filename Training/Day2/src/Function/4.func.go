package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("Demo: Math")

	a := Add(10, 20)
	b := Sub(20, 10)
	fmt.Println(a, b)

	c := Add(Add(10, 20), Sub(30, 40))

	fmt.Println(c)
}
