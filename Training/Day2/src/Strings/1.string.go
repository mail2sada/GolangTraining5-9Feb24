package main

import "fmt"

func main() {
	fmt.Println("Demo: Strings")

	var Str string = "Hello"

	fmt.Println(Str)

	var Str1 = `Hello how are you
	this is test raw string`
	fmt.Println(Str1)

	fmt.Println(Str1[0])
	fmt.Printf("%c\n", Str1[0])
}
