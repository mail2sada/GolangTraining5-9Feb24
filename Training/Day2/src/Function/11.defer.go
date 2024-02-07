package main

import "fmt"

func main() {
	fmt.Println("Demo: defer")
	defer fmt.Println("This is defered call")
	fmt.Println("This is test print")
	defer fmt.Println("This is second defer")
	fmt.Println("After the defer line")
}
