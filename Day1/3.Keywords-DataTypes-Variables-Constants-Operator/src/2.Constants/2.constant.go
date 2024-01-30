package main

import "fmt"

func main() {
	fmt.Println("Demo: Constant declaration")

	const (
		a, b int = 10, 20
		c        = "Hello"
		d, e     = true, 3.14159
	)
	fmt.Println(a, b, c, d, e)
}
