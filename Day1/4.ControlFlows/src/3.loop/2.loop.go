package main

import "fmt"

func main() {
	fmt.Println("Demo: While-Like loop")
	x := 0
	for x < 5 {
		fmt.Println(x)
		x += 1
	}
}
