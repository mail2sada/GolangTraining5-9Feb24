package main

import "fmt"

func PrintInt(x int) {
	fmt.Println("PrintInt::", x) //
}

func main() {
	fmt.Println("Demo: defer")
	i := 0
	fmt.Println(i) // 0
	defer PrintInt(i)
	i = 100
	fmt.Println(i) //100

}
