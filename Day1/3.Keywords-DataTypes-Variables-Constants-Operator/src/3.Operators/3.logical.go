package main

import "fmt"

func main() {
	fmt.Println("Demo: Logical Operators")
	p := true
	q := false
	andResult := p && q // false
	orResult := p || q  // true
	notResult := !p     // false
	fmt.Println(andResult, orResult, notResult)
}
