package main

import "fmt"

// func Square(side float64) (float64, float64) {
// 	area := side * side

// 	cir := 4 * side

// 	return area, cir
// }

func Square(side float64) (area float64, cir float64) {
	area = side * side

	cir = 4 * side
	return
}

// func Square(side float64) (area float64, cir float64) {
// 	area = side * side
// 	return
// }

func main() {
	fmt.Println("Demo: Named returns")
	a, c := Square(10)
	fmt.Println(a, c)
}
