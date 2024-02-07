package main

import "fmt"

func main() {
	a := struct {
		a, b int
		c    string
	}{a: 100, b: 200, c: "Hello"}

	b := a

	fmt.Println(a.a, b)

}
