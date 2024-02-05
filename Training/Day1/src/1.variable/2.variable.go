package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Variable Declaration")
	var test1 = 100
	fmt.Println("Value of test1 is", test1)
	fmt.Printf("Data type of test1 is %T\n", test1)

	var float = 10.10

	fmt.Println("Value of float is ", float)
	fmt.Println("Type of float is ", reflect.TypeOf(float))

	var test = true
	fmt.Println("Value of test", test)
	fmt.Printf("type of test %T", test)

	var testString = "Hello \nWorld"
	fmt.Println("testString", testString)
	fmt.Println("type of testString", reflect.TypeOf(testString))

	var myString = `This is my test 'string' ""
	whatever you like, <>`
	fmt.Println(myString)
	fmt.Println(reflect.TypeOf(myString))
}
