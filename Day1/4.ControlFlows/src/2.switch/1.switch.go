package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch case")
	weekDay := 1
	switch weekDay {
	case 1:
		fmt.Println("It's a Sunday")
	case 2:
		fmt.Println("It's a Monday")
	case 3:
		fmt.Println("It's a Tuesday")
	case 4:
		fmt.Println("It's a Wednsday")
	case 5:
		fmt.Println("It's a Thursday")
	case 6:
		fmt.Println("It's a Friday")
	case 7:
		fmt.Println("It's a Saturday")
	default:
		fmt.Println("Invalid weekday number")
	}
}
