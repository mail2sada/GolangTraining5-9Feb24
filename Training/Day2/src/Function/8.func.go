package main

import "fmt"

func main() {
	fmt.Println("Annonymous functions")

	func() {
		fmt.Println("In side annonymous function")
	}()

	x := func(welcome string) {
		fmt.Println(welcome)
	}

	x("hello")
	x("hi")
	x("test")

}
