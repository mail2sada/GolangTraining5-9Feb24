package main

import "fmt"

func main() {
	fmt.Println("Initialising slice ")

	var slice = []int{0: 100, 99: 1000, 500: 1100}

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 10000)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
