package main

import "fmt"

func main() {
	fmt.Println("Demo: Variable declaration")

	var a, b, c = 10, 20.2, false
	fmt.Println(a, b, c)
	fmt.Printf("%T %T %T", a, b, c)

}
