package main

import "fmt"

func main() {
	fmt.Println("Demo: break statement")
	x := 0
	for {
		fmt.Println(x)
		x += 1
		if x > 100 {
			break
		}
	}
}
