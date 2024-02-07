package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice len/cap")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Contents of slice", slice)

	fmt.Println("Length of slice", len(slice))

	fmt.Println("Capacity of slice", cap(slice))
}
