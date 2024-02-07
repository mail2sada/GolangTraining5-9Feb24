package main

import "fmt"

func main() {
	fmt.Println("Demo: Pointer to array")

	var array [5]int = [5]int{1, 2, 3, 4, 5}

	var ptrArray *[5]int = &array

	fmt.Println("Contents of array", array)
	fmt.Println("Contents of array using ptrArray", *ptrArray)

	fmt.Println("Contents of array using ptrArray at 0 index", ptrArray[0])

	ptr := &array

	ptr1 := &array[0]

	fmt.Println(*ptr)

	fmt.Println(*ptr1)
}
