package math

import "fmt"

func init() {
	fmt.Println("Inside init Math Pakcage sub.go")
}

func Subtract(a, b int) int {
	return Sub(a, b)
}

func Sub(a, b int) int {
	return a - b
}
