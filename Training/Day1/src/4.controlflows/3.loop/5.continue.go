package main

import "fmt"

func main() {
	fmt.Println("Demo: continue in for loop")

	i := -1

	for {
		i++
		if i%10 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 1000 && i%4 == 0 && i%101 == 0 {
			break
		}
	}
}
