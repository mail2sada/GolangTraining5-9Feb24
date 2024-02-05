package main

import "fmt"

func main() {
	fmt.Println("Demo: Arrays")
	var intArray = [5]int{0, 1, 2, 3, 4}

	for i := 0; i < 5; i++ {
		fmt.Println(intArray[i])
		intArray[i] = intArray[i] * i
		fmt.Println(intArray[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Println(intArray[i])
	}
}
