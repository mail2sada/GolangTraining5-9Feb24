package main

import "fmt"

func main() {

	fmt.Println("for loop BREAK statement")

	i := 0
	for {
		fmt.Println(i)
		if i > 10000 && i%4 == 0 && i%100 == 0 {
			break
		}
		i++
		
	}
}
