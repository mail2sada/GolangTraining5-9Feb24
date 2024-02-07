package main

import "fmt"

func TestCode(a int) {
	a = 100
	fmt.Println(a)
}

func main() {
	fmt.Println("Demi: Call by value")

	var v = 1000

	TestCode(v)

	fmt.Println(v)
}
