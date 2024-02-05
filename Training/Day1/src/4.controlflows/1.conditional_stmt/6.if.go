package main

import "fmt"

func main() {
	fmt.Println("Demo: multiple if")

	var weekDay int

	fmt.Println("Input value for weekDay")
	fmt.Scanf("%d", &weekDay)
	fmt.Println("Value of weekDay", weekDay)
	if weekDay == 1 {
		fmt.Println("Sunday")
	} else if weekDay == 2 {
		fmt.Println("Monday")
	} else {
		fmt.Println("I dont know")
	}
}
