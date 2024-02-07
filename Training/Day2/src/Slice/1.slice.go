package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice declaration")

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	fmt.Println(slice[0])
	fmt.Println(slice[5])

	slice[7] = 100

	fmt.Println(slice)

}
