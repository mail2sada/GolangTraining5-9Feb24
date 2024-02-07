package main

import "fmt"

func main() {
	fmt.Println("Demo: Delete/insert")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice)

	slice = append(slice[:6], slice[7:]...)

	fmt.Println(slice)

	slice = append(slice[:6], slice[5:]...)
	slice[6] = 100

	fmt.Println(slice)

}
