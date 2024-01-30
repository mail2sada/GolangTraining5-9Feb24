package main

import "fmt"

func main() {
	fmt.Println("Demo: Fallthrough with switch")
	day := "Thursday"
	switch day {
	case "Monday":
		fallthrough
	case "Tuesday":
		fmt.Println("It's a weekday")
	case "Wednesday":
		fmt.Println("It's hump day!")
	case "Thursday":
		fallthrough
	case "Friday":
		fmt.Println("It's almost the weekend")
	default:
		fmt.Println("It's the weekend!")
	}
}
