package main

import "fmt"

func main() {
	array := [3]interface{}{1, "Hello", "hi"}

	fmt.Println(array[:]...)
}
