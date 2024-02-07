package main

import "fmt"

func main() {
	fmt.Println("Demo: Pointers")

	var a int = 100

	var ptr *int = &a

	fmt.Println(ptr)

	fmt.Println("Value of a", a)

	fmt.Println("Value of a using pointer", *ptr)

	*ptr = 500

	fmt.Println("Value of a:", a)

	fmt.Println("Address of ptr", &ptr)
	fmt.Println("Address of a", &a)
}
