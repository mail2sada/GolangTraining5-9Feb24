package main

import "fmt"

func main() {
	fmt.Println("Demo: Constant declaration.")
	const ConstantInt16 int16 = 100
	fmt.Println("Constant Value is ", ConstantInt16)

	//ConstantInt16 = 1000 // Can't do this, as ConstantInt16 is const variable
}
