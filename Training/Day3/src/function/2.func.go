package main

import "fmt"

func TestCode(a *int) {
	*a = 100
	fmt.Println("Testcode:", *a)
}

/*
	0x000001 v = 1000 -100

	ox000008 a = 0x000001
*/

func main() {
	fmt.Println("Demo: Call by reference")

	var v = 1000

	TestCode(&v)

	fmt.Println(v)
}
