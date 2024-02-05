package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Variable Declaration")
	var a, b, c = 10, 20, 30
	fmt.Println(a, b, c)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
	fmt.Printf("%T%T%T\n", a, b, c)

	var d, e, f int8
	fmt.Println(d, f, e)
	var g float64
	fmt.Println(g)
}
