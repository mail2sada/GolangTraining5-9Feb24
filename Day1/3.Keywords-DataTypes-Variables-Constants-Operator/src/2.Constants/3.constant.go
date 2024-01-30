package main

import "fmt"

const (
	a, b int    = 10, 20
	c           = 100
	d, e        = true, 3.14159
	f           = 'h'
	g    string = "Hello"
)

func main() {
	fmt.Println("Demo: Constant global block")
	fmt.Println(a, b, c, d, e, f, g)
}
