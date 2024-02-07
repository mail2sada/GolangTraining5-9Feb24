package main

import "fmt"

func main() {
	fmt.Println("Demo: Find map")

	v := map[int]int{0: 1, 1: 2, 2: 3, 3: 4}

	_, ok := v[3]
	if ok {
		fmt.Println("Found")
	} else {
		fmt.Println("not Found")
	}

}
