package main

import "fmt"

func main() {
	fmt.Println("Demo:Switch with multiple case")
	day := "Wednesday"
	switch day {
	case "Monday", "Tuesday":
		fmt.Println("It's a weekday")
	case "Wednesday":
		fmt.Println("It's hump day!")
	case "Thursday", "Friday":
		fmt.Println("It's almost the weekend")
	default:
		fmt.Println("It's the weekend!")
	}
}
