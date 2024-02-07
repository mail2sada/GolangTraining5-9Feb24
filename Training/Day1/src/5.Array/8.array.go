package main

import "fmt"

func main() {
	fmt.Println("Demo: Length of array")
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Contents of array", array)

	fmt.Println("Length of array", len(array))
}
