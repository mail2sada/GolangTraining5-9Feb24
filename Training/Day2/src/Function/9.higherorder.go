package main

import "fmt"

func StringIterator(s string, fn func(char rune)) {
	for _, v := range s {
		fn(v)
	}
}

func main() {
	fmt.Println("Higher Order function")

	StringIterator("Hello and welcome to go training", func(char rune) {
		fmt.Println(char)
	})
}
