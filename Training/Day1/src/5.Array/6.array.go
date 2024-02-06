package main

import "fmt"

func main() {
	fmt.Println("Demo: Array filters")

	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:6]

	fmt.Println(b)

}
