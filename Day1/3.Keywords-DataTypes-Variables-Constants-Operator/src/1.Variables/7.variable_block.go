package main

import "fmt"

func main() {
	fmt.Println("Demo:Variable Block")
	var (
		a    int    = 100
		b    string = "Hello"
		c, d int    = 200, 300
		e, f        = true, 10.5
		g           = 'h'
	)

	fmt.Println(a, b, c, d, e, f, g)
}
