package main

import "fmt"

func main() {
	fmt.Println("Slice: Insert")

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:5])
	fmt.Println(slice[5:])

	slice = append(slice[:5], slice[4:]...)
	slice[5] = 100
	fmt.Println(slice)

	idx := 7

	slice = append(slice[:idx], slice[idx-1:]...)

	slice[idx] = 1000

	fmt.Println(slice)
}
