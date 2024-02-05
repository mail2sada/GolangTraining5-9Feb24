package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch case")

	weekDay := 0

	fmt.Println("Key in a value for weekday")
	fmt.Scanf("%d", &weekDay)

	switch weekDay {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("I am not sure...")
	}
}
