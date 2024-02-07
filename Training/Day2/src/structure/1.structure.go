package main

import "fmt"

type TestStruct struct {
	Integer int
	String  string
	Float   float64
	Array   [5]int32
}

func main() {
	fmt.Println("Demo: Structure")

	var testInstance TestStruct

	fmt.Println(testInstance)
	testInstance.Integer = 100
	testInstance.String = "Hello"
	testInstance.Float = 3.14159
	testInstance.Array[0] = 100
	testInstance.Array[1] = 1000
	testInstance.Array[2] = 1000
	testInstance.Array[3] = 10000
	testInstance.Array[4] = 10000
	fmt.Println(testInstance)
}
