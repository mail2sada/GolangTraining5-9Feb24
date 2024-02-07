package main

import "fmt"

func main() {
	fmt.Println("Reset Map")

	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}

	fmt.Println(m)

	for key, _ := range m {
		delete(m, key)
	}

	fmt.Println(m)

	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	fmt.Println(m)
	m = make(map[int]int)

	fmt.Println(m)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	fmt.Println(m)

	m = map[int]int{}

	m = nil

	fmt.Println(m)
}
