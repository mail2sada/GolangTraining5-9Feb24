package main

import "fmt"

var (
	a    int = 100
	b, c int = 200, 300
	d, e     = 400, 500
	f, g     = true, "Example from go training"
	h        = 'H'
)

func main() {
	fmt.Println("Demo: Variable Blok Global")
	fmt.Println(a, b, c, d, e, f, g, h)
}
