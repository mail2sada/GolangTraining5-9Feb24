package main

import "fmt"

func main() {
	fmt.Println("Demo: continue statement")
	x := -1

	for {
		x += 1
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

		if x >= 100 {
			break
		}
	}
}
