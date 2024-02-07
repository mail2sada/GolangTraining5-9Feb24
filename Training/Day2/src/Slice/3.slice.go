package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice len/cap")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Contents of slice", slice)

	fmt.Println("Length of slice", len(slice))

	fmt.Println("Capacity of slice", cap(slice))

	slice = append(slice, 11)

	fmt.Println("After append")

	fmt.Println("Contents of slice", slice)

	fmt.Println("Length of slice", len(slice))

	fmt.Println("Capacity of slice", cap(slice))

	fmt.Println("Appending elements to slice")

	for i := 0; i < 9; i++ {
		slice = append(slice, i+12)
	}

	fmt.Println("After appending additional 9 elements")

	fmt.Println("Contents of slice", slice)

	fmt.Println("Length of slice", len(slice))

	fmt.Println("Capacity of slice", cap(slice))

	fmt.Println("Appending elements to slice")

	slice = append(slice, 21)

	fmt.Println("After appending 21st element")
	fmt.Println("Contents of slice", slice)

	fmt.Println("Length of slice", len(slice))

	fmt.Println("Capacity of slice", cap(slice))

	fmt.Println("Appending elements to slice")

}
