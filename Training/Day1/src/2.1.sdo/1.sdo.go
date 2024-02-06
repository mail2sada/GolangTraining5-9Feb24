package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo sdo")

	var a, b = 10, 20

	sum := a + b
	diff := a - b
	mul := a * b
	quo := a / b
	fmt.Println(sum, diff, mul, quo)

	test := "hello"
	fmt.Println(test)

	tp := reflect.TypeOf(test)
	fmt.Println(tp)

	
}
