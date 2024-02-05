package main

import "fmt"

func main() {
	fmt.Println("Demo: Fallthrough")
	weekDay := 1

	switch weekDay {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("WeekDay")
	case 3:
		fmt.Println("Mid of the week")
	case 4, 5:
		fmt.Println("Reaching Weekend")
	case 6:
		fallthrough
	case 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid input")
	}
}
