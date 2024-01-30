package main

import "fmt"

func main() {
	fmt.Println("Demo: Bitwise Operators")
	a := 5               // 101 in binary
	b := 3               // 011 in binary
	andResult := a & b   // 1 (001 in binary)
	orResult := a | b    // 7 (111 in binary)
	xorResult := a ^ b   // 6 (110 in binary)
	leftShift := a << 1  // 10 (1010 in binary)
	rightShift := a >> 1 // 2 (10 in binary)
	bitClear := a &^ b   // 4 (100 in binary)
	fmt.Println(andResult, orResult, xorResult, leftShift, rightShift, bitClear)
}
