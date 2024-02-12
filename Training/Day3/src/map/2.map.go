package main

import "fmt"

func main() {
	fmt.Println("Demo: Maps")

	var m map[string]int = make(map[string]int)

	m["A"] = 100

	m["B"] = 1000

	m["C"] = 10000

	fmt.Println(m)

	delete(m, "A")
	fmt.Println(m)

}
