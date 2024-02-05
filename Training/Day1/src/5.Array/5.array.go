package main

import "fmt"

func main() {

	fmt.Println("Demo: Array Iterator")

	var array [5]int = [5]int{1, 2, 3, 4, 5}

	for idx, val := range array {
		fmt.Println("idx:", idx, "value:", val)
	}

	for _, val := range array {
		fmt.Println("value:", val)
	}

	for idx, _ := range array {
		fmt.Println("idx:", idx)
	}
}
