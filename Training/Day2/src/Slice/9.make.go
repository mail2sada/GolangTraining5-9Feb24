package main

import "fmt"

func main() {
	fmt.Println("Slice : make")

	slice := make([]int, 10, 100)

	fmt.Println(slice)

	fmt.Println(len(slice))

	fmt.Println(cap(slice))

	v1 := 100

	slice = append(slice, 1, 2, 3, 4, v1, 6, 7, 8, 9, 10)
	fmt.Println(slice)

	fmt.Println(len(slice))

	fmt.Println(cap(slice))

}
