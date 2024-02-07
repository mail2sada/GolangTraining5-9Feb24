package main

import "fmt"

func main() {
	fmt.Println("Demo: Empty Interface")

	var v interface{}

	v = 1
	fmt.Println(v)

	v = "Hello how are you "

	fmt.Println(v)

	v = 3.14159
	fmt.Println(v)
}
