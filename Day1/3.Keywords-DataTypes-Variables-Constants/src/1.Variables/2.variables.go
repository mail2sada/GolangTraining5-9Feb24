package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Variable declaration")
	var integer = 100
	fmt.Println("Value of interger is ", integer)
	fmt.Printf("Type of integer is %T \n", integer)
	fmt.Println("Type of integer is ", reflect.TypeOf(integer))
}
