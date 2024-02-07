package main

import "fmt"

func AddMany(elements ...int) (result int) {
	for _, val := range elements {
		result += val
	}
	return
}

func Add(a, b int, list ...int) (result int) {
	result = a + b

	for _, v := range list {
		result += v
	}
	return
}

func CustomAdd(elem ...interface{}) {
	fmt.Println(elem...)
}

func main() {
	fmt.Println("Demo: Variadic arguments")

	CustomAdd(1, "hello", 3.14159, "test")

	fmt.Println(AddMany(1, 2, 3, 4, 5, 6, 7))

	fmt.Println(AddMany())

	fmt.Println(AddMany(1))
	fmt.Println(AddMany(1, 2, 3))

	fmt.Println(Add(1, 2, 3, 4, 5, 6))
	fmt.Println(Add(1, 2))
	fmt.Println(Add(5, 6, 7))
}
