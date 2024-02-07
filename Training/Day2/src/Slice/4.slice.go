package main

import "fmt"

func main() {
	fmt.Println("Demo: nil slice")

	var slice []int

	fmt.Println(slice)

	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 0)

	fmt.Println(slice)

	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 1)

	fmt.Println(slice)

	fmt.Println(len(slice), cap(slice))

	slice = append(slice, 1)

	fmt.Println(slice)

	fmt.Println(len(slice), cap(slice))

	slice[2] = 10

	fmt.Println(slice)
}
