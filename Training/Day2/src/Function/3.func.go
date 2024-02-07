package main

import "fmt"

func Add(a int, b int) {
	fmt.Println(a + b)
}

func Subtract(a, b int) {
	fmt.Println(a - b)
}

func Multiply(a int, b int) {
	fmt.Println(a * b)
}

func main() {
	fmt.Println("Demo: Math Func")
	Add(10, 20)

	Subtract(20, 10)
	Multiply(3, 300)
}
