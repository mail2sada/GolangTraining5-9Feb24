package main

import "fmt"

func main() {
	fmt.Println("Demo: Infinite loop")
	x := 0
	for {
		fmt.Println(x)
		x += 1
	}
}
