package main

import "fmt"

type PromptedStruct struct {
	int
	string
	float64
	uint
	uint16
	uint32

	int16
	int32
	a, b int
}

func main() {
	fmt.Println("Demo: Prompted Fields")
	prs := PromptedStruct{int: 10, string: "hello", float64: 3.14159, uint: 100, a: 1000, b: 10000}
	fmt.Println(prs)
}
