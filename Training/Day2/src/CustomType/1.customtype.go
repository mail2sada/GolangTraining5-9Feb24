package main

import "fmt"

type Color uint8

const (
	Red   Color = 0
	Green Color = 1
	Blue  Color = 2
)

func main() {
	fmt.Println("Demo: Custom Type")

	var cl Color = Red

	fmt.Println(cl)
	cl = Green
	fmt.Println(cl)
}
