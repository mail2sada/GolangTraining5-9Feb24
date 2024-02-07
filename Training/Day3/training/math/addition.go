package math

import "fmt"

func init() {
	fmt.Println("Math Package addition.go init")
}

func Add(a, b int) int {
	return a + b
}

func AddMany(a, b int, elements ...int) (sum int) {
	sum = a + b

	for _, v := range elements {
		sum += v
	}
	return
}
