package main

import "fmt"

func main() {
	fmt.Println("Demo: Local Variable")

	var a = 10 // Variable is acessible across main function
	// It is scope is local to the function
	fmt.Println("Value of a is ", a)
	{
		var b = 100 // scope of b is local to block
		fmt.Println("Value of b is ", b)
	}

	//fmt.Println("Value of block variable is", b) //Uncommenting this line will lead to error
	fmt.Println("Valiue of a is", a) // As a is local to a function it can be accessed across function
}
