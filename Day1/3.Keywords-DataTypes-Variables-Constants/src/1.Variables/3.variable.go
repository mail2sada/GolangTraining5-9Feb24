package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo:Variable declaration")
	var a, b, c = 10, 20, 30
	fmt.Println(a, b, c)
	fmt.Println("Type of a is ", reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))

	var d, e, f = 1, 1.1, true
	fmt.Println(d, e, f)
	fmt.Printf("Type of d %T, type of e %T, type of f %T", d, e, f)
}
