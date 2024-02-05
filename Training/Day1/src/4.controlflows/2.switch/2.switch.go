package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch")

	weekDay := -2

	switch weekDay {
	case 1, 2, 3, 4, 5:
		fmt.Println("WeekDay")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid value for weekDay")
	}
}
