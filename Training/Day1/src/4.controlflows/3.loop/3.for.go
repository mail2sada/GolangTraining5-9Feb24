package main

import "fmt"

func main() {
	fmt.Println("Demo:Infinite loops")

	i := 0

	for {
		fmt.Println(i)
		i++
	}
}
