package main

import "fmt"

func TestCode(a int) {
	a = 100
	fmt.Println(a)
}

//  0x0001 -> v 1000
//
// 0x0008 -> a 1000/ 100
//
//

func main() {
	fmt.Println("Demi: Call by value")

	var v = 1000

	TestCode(v)

	fmt.Println(v)
}
