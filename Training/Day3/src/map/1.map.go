package main

import "fmt"

func main() {
	fmt.Println("Demo: Maps")

	mmap := map[string]int{"A": 1, "B": 2, "C": 3}

	fmt.Println(mmap)

	i := mmap["A"]

	fmt.Println(i)

	mmap["X"] = 100

	fmt.Println(mmap)
}
