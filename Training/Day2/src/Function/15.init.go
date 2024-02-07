package main

import "fmt"

func init() {
	fmt.Println("Inside init")
}

func init() {
	fmt.Println("Inside 2nd init")
}

func init() {
	fmt.Println("We are 3rd init")
}

func main() {
	fmt.Println("Demo : init")
	fmt.Println("We are in main")
}

func init() {
	fmt.Println("After main::we are in 4th init")
}
